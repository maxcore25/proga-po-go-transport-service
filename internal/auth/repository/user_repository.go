package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uuid.UUID) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Find(filter dto.UserFilter) ([]*model.User, error)
	UpdateByID(id uuid.UUID, updateData dto.UpdateUserRequest) error
	DeleteByID(id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uuid.UUID) (*model.User, error) {
	var u model.User
	if err := r.db.First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var u model.User
	if err := r.db.First(&u, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Find(filter dto.UserFilter) ([]*model.User, error) {
	db := r.db.Model(&model.User{})

	if filter.Role != nil {
		db = db.Where("role = ?", *filter.Role)
	}

	if filter.KnowledgeLevel != nil {
		db = db.Where("knowledge_level = ?", *filter.KnowledgeLevel)
	}

	var users []*model.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) UpdateByID(id uuid.UUID, updateData dto.UpdateUserRequest) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *userRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.User{}).Error
}
