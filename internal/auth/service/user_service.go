package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/model"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/repository"
	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/utils"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetUser(id uuid.UUID) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	GetUsers(filter dto.UserFilter) ([]*model.User, error)
	UpdateUserByID(id uuid.UUID, updateData dto.UpdateUserRequest) error
	DeleteUserByID(id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (*model.User, error) {
	// 1. Check if email already exists
	existingUser, err := s.repo.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// 2. Hash password before storing
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	var middleNamePtr *string
	if req.MiddleName != "" {
		middleNamePtr = &req.MiddleName
	}

	var phonePtr *string
	if req.Phone != "" {
		phonePtr = &req.Phone
	}

	user := &model.User{
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		MiddleName:        middleNamePtr,
		Email:             req.Email,
		Password:          hashedPassword,
		Phone:             phonePtr,
		KnowledgeLevel:    model.KnowledgeLevel(req.KnowledgeLevel),
		Role:              model.UserRole(req.Role),
		Portfolio:         req.Portfolio,
		Rating:            req.Rating,
		TestimonialsCount: req.TestimonialsCount,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByEmail(email string) (*model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *userService) GetUser(id uuid.UUID) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]*model.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUsers(filter dto.UserFilter) ([]*model.User, error) {
	return s.repo.Find(filter)
}

func (s *userService) UpdateUserByID(id uuid.UUID, updateData dto.UpdateUserRequest) error {
	return s.repo.UpdateByID(id, updateData)
}

func (s *userService) DeleteUserByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
