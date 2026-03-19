package dto

import "time"

// CreateCardRequest represents a card creation request.
// @Description Card creation request payload
// @Name CreateCardRequest
type CreateCardRequest struct {
	CardNumber  string     `json:"cardNumber" binding:"required,min=4,max=64" example:"A1B2C3D4"`
	OwnerName   string     `json:"ownerName" binding:"required,min=2,max=100" example:"Ivan Ivanov"`
	Balance     int        `json:"balance" example:"5000"`
	KeyID       string     `json:"keyId" binding:"required,uuid" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty" example:"2027-12-31T23:59:59Z"`
	IsBlocked   bool       `json:"isBlocked" example:"false"`
	BlockReason *string    `json:"blockReason,omitempty" example:"manual block"`
}
