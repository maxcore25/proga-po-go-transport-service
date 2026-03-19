package dto

import "time"

// PaymentAuthResponse — ответ сервера терминалу после проверки транзакции.
//
// Терминал должен показать пассажиру зелёный сигнал если Approved=true,
// красный — если Approved=false. Message содержит машиночитаемую причину отказа.
//
// @Description Terminal payment authorization response
// @Name PaymentAuthResponse
type PaymentAuthResponse struct {
	// Результат авторизации
	Approved bool `json:"approved" example:"true"`
	// Уникальный ID созданной транзакции (только при Approved=true)
	TransactionID *string `json:"transactionId,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
	// Машиночитаемый код результата
	// Возможные значения: "approved" | "card_not_found" | "card_blocked" |
	// "card_expired" | "insufficient_funds" | "terminal_not_found" | "terminal_inactive"
	Code string `json:"code" example:"approved"`
	// Человекочитаемое сообщение
	Message string `json:"message" example:"Payment authorized successfully"`
	// Баланс после списания в копейках (только при Approved=true)
	BalanceAfter *int `json:"balanceAfter,omitempty" example:"96500"`
	// Время транзакции (UTC)
	ProcessedAt time.Time `json:"processedAt"`
}

// Коды результатов авторизации — используются в Code поле ответа
const (
	CodeApproved          = "approved"
	CodeCardNotFound      = "card_not_found"
	CodeCardBlocked       = "card_blocked"
	CodeCardExpired       = "card_expired"
	CodeInsufficientFunds = "insufficient_funds"
	CodeTerminalNotFound  = "terminal_not_found"
	CodeTerminalInactive  = "terminal_inactive"
)
