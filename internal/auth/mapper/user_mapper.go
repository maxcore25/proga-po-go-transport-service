package mapper

import (
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/model"
)

func NewUserResponse(u *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		FullName:  u.FullName,
		IsAdmin:   u.IsAdmin,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
