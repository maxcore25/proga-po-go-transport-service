package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	cardRepo "github.com/maxcore25/proga-po-go-transport-service/internal/cards/repository"
	keyRepo "github.com/maxcore25/proga-po-go-transport-service/internal/keys/repository"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminal_api/dto"
	terminalRepo "github.com/maxcore25/proga-po-go-transport-service/internal/terminals/repository"
	transactionModel "github.com/maxcore25/proga-po-go-transport-service/internal/transactions/model"
	transactionRepo "github.com/maxcore25/proga-po-go-transport-service/internal/transactions/repository"
)

// TerminalAPIService — бизнес-логика терминального API.
// Не требует JWT: терминалы идентифицируются по serial_number.
type TerminalAPIService interface {
	// AuthorizePayment выполняет все проверки и при успехе списывает баланс.
	AuthorizePayment(req dto.PaymentAuthRequest) (*dto.PaymentAuthResponse, error)
	// LoadKeys возвращает все активные ключи шифрования MIFARE.
	LoadKeys(terminalSerial string) (*dto.KeysLoadResponse, error)
}

type terminalAPIService struct {
	cardRepo        cardRepo.CardRepository
	terminalRepo    terminalRepo.TerminalRepository
	transactionRepo transactionRepo.TransactionRepository
	keyRepo         keyRepo.KeyRepository
}

func NewTerminalAPIService(
	cardRepo cardRepo.CardRepository,
	terminalRepo terminalRepo.TerminalRepository,
	transactionRepo transactionRepo.TransactionRepository,
	keyRepo keyRepo.KeyRepository,
) TerminalAPIService {
	return &terminalAPIService{
		cardRepo:        cardRepo,
		terminalRepo:    terminalRepo,
		transactionRepo: transactionRepo,
		keyRepo:         keyRepo,
	}
}

// AuthorizePayment — ядро системы. Порядок проверок строго соблюдается:
//
//  1. Терминал существует и активен
//  2. Карта существует
//  3. Карта не заблокирована
//  4. Карта не просрочена
//  5. Баланс достаточен
//
// Только после всех проверок — списываем баланс и записываем транзакцию.
// Declined-транзакции тоже записываются (для аудита).
func (s *terminalAPIService) AuthorizePayment(req dto.PaymentAuthRequest) (*dto.PaymentAuthResponse, error) {
	now := time.Now().UTC()

	// ── Шаг 1: Проверяем терминал ──────────────────────────────────────────
	terminal, err := s.terminalRepo.GetBySerialNumber(req.TerminalSerial)
	if err != nil {
		return declined(dto.CodeTerminalNotFound, "Terminal not found", now), nil
	}
	if !terminal.IsActive {
		return declined(dto.CodeTerminalInactive, "Terminal is deactivated", now), nil
	}

	// Обновляем last_seen_at терминала (не блокируем при ошибке)
	_ = s.terminalRepo.UpdateLastSeen(terminal.ID, now)

	// ── Шаг 2: Ищем карту по UID ───────────────────────────────────────────
	card, err := s.cardRepo.GetByCardNumber(req.CardNumber)
	if err != nil {
		// Declined-транзакцию не записываем — нет card_id для FK
		return declined(dto.CodeCardNotFound, "Card not found", now), nil
	}

	// Все последующие отказы записываем в transactions для аудита
	// ── Шаг 3: Карта заблокирована? ────────────────────────────────────────
	if card.IsBlocked {
		reason := "Card is blocked"
		if card.BlockReason != nil {
			reason = fmt.Sprintf("Card is blocked: %s", *card.BlockReason)
		}
		_ = s.recordDeclined(card.ID, terminal.ID, req.Amount, card.Balance, dto.CodeCardBlocked, reason, now)
		return declined(dto.CodeCardBlocked, reason, now), nil
	}

	// ── Шаг 4: Карта просрочена? ────────────────────────────────────────────
	if card.ExpiresAt != nil && card.ExpiresAt.Before(now) {
		_ = s.recordDeclined(card.ID, terminal.ID, req.Amount, card.Balance, dto.CodeCardExpired, "Card has expired", now)
		return declined(dto.CodeCardExpired, "Card has expired", now), nil
	}

	// ── Шаг 5: Достаточно средств? ──────────────────────────────────────────
	if card.Balance < req.Amount {
		msg := fmt.Sprintf("Insufficient funds: balance %d, required %d", card.Balance, req.Amount)
		_ = s.recordDeclined(card.ID, terminal.ID, req.Amount, card.Balance, dto.CodeInsufficientFunds, msg, now)
		return declined(dto.CodeInsufficientFunds, "Insufficient funds", now), nil
	}

	// ── Всё ОК: списываем баланс и записываем транзакцию ───────────────────
	balanceBefore := card.Balance
	balanceAfter := card.Balance - req.Amount

	if err := s.cardRepo.UpdateBalance(card.ID, balanceAfter); err != nil {
		return nil, fmt.Errorf("update balance: %w", err)
	}

	txID := uuid.New()
	txIDStr := txID.String()
	declineReason := (*string)(nil)

	tx := &transactionModel.Transaction{
		ID:            txID,
		CardID:        card.ID,
		TerminalID:    terminal.ID,
		Amount:        req.Amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balanceAfter,
		Status:        "approved",
		DeclineReason: declineReason,
		CreatedAt:     now,
	}

	if err := s.transactionRepo.Create(tx); err != nil {
		// Критично: баланс уже списан, но транзакция не записана.
		// Откатываем баланс — атомарность без SQL-транзакции (SQLite WAL).
		_ = s.cardRepo.UpdateBalance(card.ID, balanceBefore)
		return nil, fmt.Errorf("record transaction: %w", err)
	}

	return &dto.PaymentAuthResponse{
		Approved:      true,
		TransactionID: &txIDStr,
		Code:          dto.CodeApproved,
		Message:       "Payment authorized successfully",
		BalanceAfter:  &balanceAfter,
		ProcessedAt:   now,
	}, nil
}

// LoadKeys — отдаёт терминалу все активные ключи MIFARE.
//
// Терминал идентифицируется по serial_number и должен быть активен.
// key_value передаётся в hex — терминал использует его для расшифровки секторов карты.
func (s *terminalAPIService) LoadKeys(terminalSerial string) (*dto.KeysLoadResponse, error) {
	// Проверяем что терминал знаком системе и активен
	terminal, err := s.terminalRepo.GetBySerialNumber(terminalSerial)
	if err != nil {
		return nil, errors.New("terminal not found")
	}
	if !terminal.IsActive {
		return nil, errors.New("terminal is deactivated")
	}

	_ = s.terminalRepo.UpdateLastSeen(terminal.ID, time.Now().UTC())

	keys, err := s.keyRepo.GetAllActive()
	if err != nil {
		return nil, fmt.Errorf("load keys: %w", err)
	}

	result := make([]dto.KeyLoadResponse, 0, len(keys))
	for _, k := range keys {
		result = append(result, dto.KeyLoadResponse{
			ID:       k.ID.String(),
			Name:     k.Name,
			KeyValue: k.KeyValue,
			KeyType:  k.KeyType,
			Sector:   k.Sector,
		})
	}

	return &dto.KeysLoadResponse{
		Keys:     result,
		Count:    len(result),
		IssuedAt: time.Now().UTC(),
	}, nil
}

// ── Вспомогательные функции ──────────────────────────────────────────────────

// declined строит отказной ответ без записи транзакции.
func declined(code, message string, at time.Time) *dto.PaymentAuthResponse {
	return &dto.PaymentAuthResponse{
		Approved:    false,
		Code:        code,
		Message:     message,
		ProcessedAt: at,
	}
}

// recordDeclined создаёт declined-транзакцию для аудита.
// Не возвращает ошибку наружу — отказ пишем best-effort.
func (s *terminalAPIService) recordDeclined(
	cardID, terminalID uuid.UUID,
	amount, balanceBefore int,
	code, reason string,
	at time.Time,
) error {
	tx := &transactionModel.Transaction{
		ID:            uuid.New(),
		CardID:        cardID,
		TerminalID:    terminalID,
		Amount:        amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balanceBefore, // баланс не изменился
		Status:        "declined",
		DeclineReason: &reason,
		CreatedAt:     at,
	}
	return s.transactionRepo.Create(tx)
}
