package dto

import (
	"time"

	"github.com/google/uuid"
)

// TransactionResponse represents the structure of a transaction returned in API responses.
// @Description Transaction response payload.
// @Name TransactionResponse
type TransactionResponse struct {
	ID            uuid.UUID `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	CardID        uuid.UUID `json:"cardId" example:"b8098c1a-f86e-11da-bd1a-00112444be1e"`
	TerminalID    uuid.UUID `json:"terminalId" example:"c8098c1a-f86e-11da-bd1a-00112444be1e"`
	Amount        int       `json:"amount" example:"2500"`
	BalanceBefore int       `json:"balanceBefore" example:"10000"`
	BalanceAfter  int       `json:"balanceAfter" example:"7500"`
	Status        string    `json:"status" example:"approved"`
	DeclineReason *string   `json:"declineReason,omitempty" example:"insufficient funds"`
	CreatedAt     time.Time `json:"createdAt" example:"2025-11-12T19:45:00Z"`
}
