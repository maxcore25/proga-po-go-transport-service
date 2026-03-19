package dto

// CreateUserRequest represents a user registration request.
// @Description User registration request payload
// @Name CreateUserRequest
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=2,max=50" example:"ivanov"`
	FullName string `json:"fullName" binding:"required,min=2,max=50" example:"Иван Иванов"`
	Password string `json:"password" binding:"required,min=6" example:"qwe123"`
	IsAdmin  bool   `json:"isAdmin" example:"false"`
}
