package dto

import "time"

// UpdateCardRequest represents a card update request.
// @Description Card update request payload
// @Name UpdateCardRequest
type UpdateCardRequest struct {
	CardNumber  *string    `json:"cardNumber,omitempty" example:"A1B2C3D4"`
	OwnerName   *string    `json:"ownerName,omitempty" example:"Ivan Ivanov"`
	Balance     *int       `json:"balance,omitempty" example:"5000"`
	IsBlocked   *bool      `json:"isBlocked,omitempty" example:"false"`
	BlockReason *string    `json:"blockReason,omitempty" example:"manual block"`
	KeyID       *string    `json:"keyId,omitempty" binding:"omitempty,uuid" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty" example:"2027-12-31T23:59:59Z"`
}
