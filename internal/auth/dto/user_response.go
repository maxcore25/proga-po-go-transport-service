package dto

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse represents the structure of a user returned in API responses.
// @Description User response payload.
// @Name UserResponse
type UserResponse struct {
	ID             uuid.UUID `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	FirstName      string    `json:"firstName" example:"Иван"`
	LastName       string    `json:"lastName" example:"Иванов"`
	MiddleName     *string   `json:"middleName,omitempty" example:"Иванович"`
	Email          string    `json:"email" example:"user@mail.ru"`
	Phone          *string   `json:"phone,omitempty" example:"+77010000000"`
	KnowledgeLevel string    `json:"knowledgeLevel" example:"beginner"`
	Role           string    `json:"role" example:"tutor"`

	// Tutor-specific fields (optional)
	Rating            *float64 `json:"rating,omitempty" example:"4.8"`
	Portfolio         *string  `json:"portfolio,omitempty" example:"Experienced backend developer with 5+ years in Go."`
	TestimonialsCount *int     `json:"testimonialsCount,omitempty" example:"12"`

	CreatedAt time.Time `json:"createdAt" example:"2025-11-12T19:45:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2025-11-12T19:45:00Z"`
}
