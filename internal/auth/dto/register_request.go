package dto

// RegisterRequest represents the payload for user registration.
// @Description Register request payload
// @Name RegisterRequest
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=2,max=50" example:"Иван"`
	FullName string `json:"fullName" binding:"required,min=2,max=50" example:"Иван Иванов"`
	Password string `json:"password" binding:"required,min=6,max=100" example:"qwe123"`
}
