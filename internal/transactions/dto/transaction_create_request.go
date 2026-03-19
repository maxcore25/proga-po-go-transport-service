package dto

// CreateTransactionRequest represents a transaction creation request.
// @Description Transaction creation request payload
// @Name CreateTransactionRequest
type CreateTransactionRequest struct {
	CardID        string  `json:"cardId" binding:"required,uuid" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	TerminalID    string  `json:"terminalId" binding:"required,uuid" example:"b8098c1a-f86e-11da-bd1a-00112444be1e"`
	Amount        int     `json:"amount" binding:"required" example:"2500"`
	BalanceBefore int     `json:"balanceBefore" binding:"required" example:"10000"`
	BalanceAfter  int     `json:"balanceAfter" binding:"required" example:"7500"`
	Status        string  `json:"status" binding:"required,oneof=approved declined" example:"approved"`
	DeclineReason *string `json:"declineReason,omitempty" example:"insufficient funds"`
}
