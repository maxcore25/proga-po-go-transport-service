package dto

import "time"

// UpdateTerminalRequest represents a terminal update request.
// @Description Terminal update request payload
// @Name UpdateTerminalRequest
type UpdateTerminalRequest struct {
	SerialNumber *string    `json:"serialNumber,omitempty" example:"TERM-0001"`
	Name         *string    `json:"name,omitempty" example:"Bus 42 Front"`
	Location     *string    `json:"location,omitempty" example:"Main Depot"`
	Route        *string    `json:"route,omitempty" example:"42"`
	IsActive     *bool      `json:"isActive,omitempty" example:"true"`
	LastSeenAt   *time.Time `json:"lastSeenAt,omitempty" example:"2026-03-19T12:00:00Z"`
}
