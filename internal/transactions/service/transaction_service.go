package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/model"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/repository"
)

type TransactionService interface {
	CreateTransaction(req dto.CreateTransactionRequest) (*model.Transaction, error)
	GetTransaction(id uuid.UUID) (*model.Transaction, error)
	GetAllTransactions() ([]*model.Transaction, error)
	GetTransactions(filter dto.TransactionFilter) ([]*model.Transaction, error)
	UpdateTransactionByID(id uuid.UUID, updateData dto.UpdateTransactionRequest) error
	DeleteTransactionByID(id uuid.UUID) error
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return &transactionService{repo: r}
}

func (s *transactionService) CreateTransaction(req dto.CreateTransactionRequest) (*model.Transaction, error) {
	cardID, err := uuid.Parse(req.CardID)
	if err != nil {
		return nil, errors.New("invalid card_id")
	}
	terminalID, err := uuid.Parse(req.TerminalID)
	if err != nil {
		return nil, errors.New("invalid terminal_id")
	}

	transaction := &model.Transaction{
		ID:            uuid.New(),
		CardID:        cardID,
		TerminalID:    terminalID,
		Amount:        req.Amount,
		BalanceBefore: req.BalanceBefore,
		BalanceAfter:  req.BalanceAfter,
		Status:        req.Status,
		DeclineReason: req.DeclineReason,
	}

	if err := s.repo.Create(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) GetTransaction(id uuid.UUID) (*model.Transaction, error) {
	transaction, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("transaction not found")
	}
	return transaction, nil
}

func (s *transactionService) GetAllTransactions() ([]*model.Transaction, error) {
	transactions, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *transactionService) GetTransactions(filter dto.TransactionFilter) ([]*model.Transaction, error) {
	return s.repo.Find(filter)
}

func (s *transactionService) UpdateTransactionByID(id uuid.UUID, updateData dto.UpdateTransactionRequest) error {
	return s.repo.UpdateByID(id, updateData)
}

func (s *transactionService) DeleteTransactionByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
