package dto

import "github.com/google/uuid"

// CardFilter represents query parameters for filtering cards.
// swagger:parameters listCards
type CardFilter struct {
	// IsBlocked is an optional filter by blocked status.
	// in: query
	IsBlocked *bool `form:"is_blocked" json:"is_blocked"`

	// KeyID is an optional filter by key ID.
	// in: query
	KeyID *uuid.UUID `form:"key_id" json:"key_id"`
}
