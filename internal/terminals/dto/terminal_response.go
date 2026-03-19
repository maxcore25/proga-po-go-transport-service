package dto

import (
	"time"

	"github.com/google/uuid"
)

// TerminalResponse represents the structure of a terminal returned in API responses.
// @Description Terminal response payload.
// @Name TerminalResponse
type TerminalResponse struct {
	ID           uuid.UUID  `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	SerialNumber string     `json:"serialNumber" example:"TERM-0001"`
	Name         string     `json:"name" example:"Bus 42 Front"`
	Location     string     `json:"location" example:"Main Depot"`
	Route        *string    `json:"route,omitempty" example:"42"`
	IsActive     bool       `json:"isActive" example:"true"`
	LastSeenAt   *time.Time `json:"lastSeenAt,omitempty" example:"2026-03-19T12:00:00Z"`
	CreatedAt    time.Time  `json:"createdAt" example:"2025-11-12T19:45:00Z"`
	UpdatedAt    time.Time  `json:"updatedAt" example:"2025-11-12T19:45:00Z"`
}
