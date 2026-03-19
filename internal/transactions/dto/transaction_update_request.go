package dto

// UpdateTransactionRequest represents a transaction update request.
// @Description Transaction update request payload
// @Name UpdateTransactionRequest
type UpdateTransactionRequest struct {
	CardID        *string `json:"cardId,omitempty" binding:"omitempty,uuid" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	TerminalID    *string `json:"terminalId,omitempty" binding:"omitempty,uuid" example:"b8098c1a-f86e-11da-bd1a-00112444be1e"`
	Amount        *int    `json:"amount,omitempty" example:"2500"`
	BalanceBefore *int    `json:"balanceBefore,omitempty" example:"10000"`
	BalanceAfter  *int    `json:"balanceAfter,omitempty" example:"7500"`
	Status        *string `json:"status,omitempty" binding:"omitempty,oneof=approved declined" example:"approved"`
	DeclineReason *string `json:"declineReason,omitempty" example:"insufficient funds"`
}
