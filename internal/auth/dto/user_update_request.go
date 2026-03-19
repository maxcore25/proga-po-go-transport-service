package dto

// UpdateUserRequest represents a user update request.
// @Description User update request payload
// @Name UpdateUserRequest
type UpdateUserRequest struct {
	Username *string `json:"username,omitempty" example:"ivan"`
	FullName *string `json:"fullName,omitempty" example:"Иван Иванов"`
	IsAdmin  *bool   `json:"isAdmin,omitempty" example:"false"`
	IsActive *bool   `json:"isActive,omitempty" example:"true"`
}
