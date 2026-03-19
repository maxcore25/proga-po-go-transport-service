package dto

import "time"

// CreateTerminalRequest represents a terminal creation request.
// @Description Terminal creation request payload
// @Name CreateTerminalRequest
type CreateTerminalRequest struct {
	SerialNumber string     `json:"serialNumber" binding:"required,min=2,max=100" example:"TERM-0001"`
	Name         string     `json:"name" binding:"required,min=2,max=100" example:"Bus 42 Front"`
	Location     string     `json:"location" binding:"required,min=2,max=200" example:"Main Depot"`
	Route        *string    `json:"route,omitempty" example:"42"`
	IsActive     bool       `json:"isActive" example:"true"`
	LastSeenAt   *time.Time `json:"lastSeenAt,omitempty" example:"2026-03-19T12:00:00Z"`
}
