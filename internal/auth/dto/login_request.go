package dto

// LoginRequest represents user login data.
// @Description Login request payload
// @Name LoginRequest
type LoginRequest struct {
	Username string `json:"username" binding:"required,min=2,max=50" example:"Иван"`
	Password string `json:"password" binding:"required" example:"qwe123"`
}
