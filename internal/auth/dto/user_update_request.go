package dto

// UpdateUserRequest represents a user update request.
// @Description User update request payload
// @Name UpdateUserRequest
type UpdateUserRequest struct {
	FirstName      *string `json:"firstName,omitempty" example:"Иван"`
	LastName       *string `json:"lastName,omitempty" example:"Иванов"`
	MiddleName     *string `json:"middleName,omitempty" example:"Иванович"`
	Email          *string `json:"email,omitempty" example:"user@mail.ru"`
	Phone          *string `json:"phone,omitempty" example:"+77010000000"`
	KnowledgeLevel *string `json:"knowledgeLevel,omitempty" example:"beginner"`
	Role           *string `json:"role,omitempty" example:"client"`

	Rating            *float64 `json:"rating,omitempty" example:"4.8"`
	Portfolio         *string  `json:"portfolio,omitempty" example:"https://myportfolio.com"`
	TestimonialsCount *int     `json:"testimonialsCount,omitempty" example:"15"`
}
