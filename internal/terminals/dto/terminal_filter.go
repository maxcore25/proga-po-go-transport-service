package dto

// TerminalFilter represents query parameters for filtering terminals.
// swagger:parameters listTerminals
type TerminalFilter struct {
	// IsActive is an optional filter by active status.
	// in: query
	IsActive *bool `form:"is_active" json:"is_active"`
}
