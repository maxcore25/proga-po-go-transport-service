package dto

// CardFilter represents query parameters for filtering cards.
// swagger:parameters listCards
type CardFilter struct {
	// IsBlocked is an optional filter by blocked status.
	// in: query
	IsBlocked *bool `form:"is_blocked" json:"is_blocked"`
}
