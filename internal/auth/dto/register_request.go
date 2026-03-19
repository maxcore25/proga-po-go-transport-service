package dto

// RegisterRequest represents the payload for user registration.
// @Description Register request payload
// @Name RegisterRequest
type RegisterRequest struct {
	FirstName      string  `json:"firstName" binding:"required,min=2,max=50" example:"Иван"`
	LastName       string  `json:"lastName" binding:"required,min=2,max=50" example:"Иванов"`
	MiddleName     *string `json:"middleName,omitempty" example:"Иванович"`
	Email          string  `json:"email" binding:"required,email" example:"user@mail.ru"`
	Password       string  `json:"password" binding:"required,min=6,max=100" example:"qwe123"`
	Phone          *string `json:"phone,omitempty" example:"+77010000000"`
	KnowledgeLevel string  `json:"knowledgeLevel" binding:"required,oneof=beginner intermediate advanced" example:"beginner"`
}
