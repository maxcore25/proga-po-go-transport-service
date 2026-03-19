package dto

// KeyFilter represents query parameters for filtering keys.
// swagger:parameters listKeys
type KeyFilter struct {
	// IsActive is an optional filter by active status.
	// in: query
	IsActive *bool `form:"is_active" json:"is_active"`

	// KeyType is an optional filter for key type (true/false).
	// If your KeyType is some domain-specific boolean, adjust this comment.
	// in: query
	KeyType *bool `form:"key_type" json:"key_type"`
}
