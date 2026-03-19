package dto

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse represents the structure of a user returned in API responses.
// @Description User response payload.
// @Name UserResponse
type UserResponse struct {
	ID        uuid.UUID `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	Username  string    `json:"username" example:"Иван"`
	FullName  string    `json:"fullName" example:"Иван Иванов"`
	IsAdmin   bool      `json:"isAdmin" example:"false"`
	IsActive  bool      `json:"isActive" example:"true"`
	CreatedAt time.Time `json:"createdAt" example:"2025-11-12T19:45:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2025-11-12T19:45:00Z"`
}
