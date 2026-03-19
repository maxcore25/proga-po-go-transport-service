package dto

import (
	"time"

	"github.com/google/uuid"
)

// CardResponse represents the structure of a card returned in API responses.
// @Description Card response payload.
// @Name CardResponse
type CardResponse struct {
	ID          uuid.UUID  `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	CardNumber  string     `json:"cardNumber" example:"A1B2C3D4"`
	OwnerName   string     `json:"ownerName" example:"Ivan Ivanov"`
	Balance     int        `json:"balance" example:"5000"`
	IsBlocked   bool       `json:"isBlocked" example:"false"`
	BlockReason *string    `json:"blockReason,omitempty" example:"manual block"`
	KeyID       uuid.UUID  `json:"keyId" example:"b8098c1a-f86e-11da-bd1a-00112444be1e"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty" example:"2027-12-31T23:59:59Z"`
	CreatedAt   time.Time  `json:"createdAt" example:"2025-11-12T19:45:00Z"`
	UpdatedAt   time.Time  `json:"updatedAt" example:"2025-11-12T19:45:00Z"`
}
