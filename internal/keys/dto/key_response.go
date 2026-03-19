package dto

import (
	"time"

	"github.com/google/uuid"
)

// KeyResponse represents the structure of a key returned in API responses.
// @Description Key response payload.
// @Name KeyResponse
type KeyResponse struct {
	ID          uuid.UUID `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	Name        string    `json:"name" example:"Main Depot Key A"`
	KeyValue    string    `json:"keyValue" example:"A1B2C3D4E5F6"`
	KeyType     string    `json:"keyType" example:"A"`
	Sector      int       `json:"sector" example:"0"`
	Description *string   `json:"description,omitempty" example:"Primary key for buses"`
	IsActive    bool      `json:"isActive" example:"true"`
	CreatedAt   time.Time `json:"createdAt" example:"2025-11-12T19:45:00Z"`
	UpdatedAt   time.Time `json:"updatedAt" example:"2025-11-12T19:45:00Z"`
}
