package dto

// KeyFilter represents query parameters for filtering keys.
// swagger:parameters listKeys
type KeyFilter struct {
	// IsActive is an optional filter by active status.
	// in: query
	IsActive *bool `form:"is_active" json:"is_active"`
}
