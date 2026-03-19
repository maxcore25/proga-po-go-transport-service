package dto

// UserFilter represents query parameters for filtering users.
// swagger:parameters listUsers
type UserFilter struct {
	// Role is an optional filter by user role.
	// in: query
	Role *string `form:"role" json:"role"`
	// KnowledgeLevel is an optional filter by knowledge level.
	// in: query
	KnowledgeLevel *string `form:"knowledge_level" json:"knowledge_level"`
}
