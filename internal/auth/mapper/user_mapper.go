package mapper

import (
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/model"
)

func NewUserResponse(u *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:                u.ID,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		MiddleName:        u.MiddleName,
		Email:             u.Email,
		Phone:             u.Phone,
		KnowledgeLevel:    string(u.KnowledgeLevel),
		Role:              string(u.Role),
		Rating:            u.Rating,
		Portfolio:         u.Portfolio,
		TestimonialsCount: u.TestimonialsCount,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
	}
}
