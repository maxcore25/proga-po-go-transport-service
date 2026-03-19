package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/model"
	"gorm.io/gorm"
)

type TerminalRepository interface {
	Create(terminal *model.Terminal) error
	GetByID(id uuid.UUID) (*model.Terminal, error)
	GetBySerialNumber(serialNumber string) (*model.Terminal, error)
	GetAll() ([]*model.Terminal, error)
	Find(filter dto.TerminalFilter) ([]*model.Terminal, error)
	UpdateByID(id uuid.UUID, updateData dto.UpdateTerminalRequest) error
	DeleteByID(id uuid.UUID) error
}

type terminalRepository struct {
	db *gorm.DB
}

func NewTerminalRepository(db *gorm.DB) TerminalRepository {
	return &terminalRepository{db: db}
}

func (r *terminalRepository) Create(terminal *model.Terminal) error {
	return r.db.Create(terminal).Error
}

func (r *terminalRepository) GetByID(id uuid.UUID) (*model.Terminal, error) {
	var t model.Terminal
	if err := r.db.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *terminalRepository) GetBySerialNumber(serialNumber string) (*model.Terminal, error) {
	var t model.Terminal
	if err := r.db.First(&t, "serial_number = ?", serialNumber).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *terminalRepository) GetAll() ([]*model.Terminal, error) {
	var terminals []*model.Terminal
	if err := r.db.Find(&terminals).Error; err != nil {
		return nil, err
	}
	return terminals, nil
}

func (r *terminalRepository) Find(filter dto.TerminalFilter) ([]*model.Terminal, error) {
	db := r.db.Model(&model.Terminal{})

	if filter.IsActive != nil {
		db = db.Where("is_active = ?", *filter.IsActive)
	}

	var terminals []*model.Terminal
	if err := db.Find(&terminals).Error; err != nil {
		return nil, err
	}

	return terminals, nil
}

func (r *terminalRepository) UpdateByID(id uuid.UUID, updateData dto.UpdateTerminalRequest) error {
	return r.db.Model(&model.Terminal{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *terminalRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.Terminal{}).Error
}
