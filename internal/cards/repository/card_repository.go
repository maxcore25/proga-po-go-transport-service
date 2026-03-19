package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/model"
	"gorm.io/gorm"
)

type CardRepository interface {
	Create(card *model.Card) error
	GetByID(id uuid.UUID) (*model.Card, error)
	GetByCardNumber(cardNumber string) (*model.Card, error)
	GetAll() ([]*model.Card, error)
	Find(filter dto.CardFilter) ([]*model.Card, error)
	UpdateByID(id uuid.UUID, updateData dto.UpdateCardRequest) error
	DeleteByID(id uuid.UUID) error
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) Create(card *model.Card) error {
	return r.db.Create(card).Error
}

func (r *cardRepository) GetByID(id uuid.UUID) (*model.Card, error) {
	var c model.Card
	if err := r.db.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *cardRepository) GetByCardNumber(cardNumber string) (*model.Card, error) {
	var c model.Card
	if err := r.db.First(&c, "card_number = ?", cardNumber).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *cardRepository) GetAll() ([]*model.Card, error) {
	var cards []*model.Card
	if err := r.db.Find(&cards).Error; err != nil {
		return nil, err
	}
	return cards, nil
}

func (r *cardRepository) Find(filter dto.CardFilter) ([]*model.Card, error) {
	db := r.db.Model(&model.Card{})

	if filter.IsBlocked != nil {
		db = db.Where("is_blocked = ?", *filter.IsBlocked)
	}

	var cards []*model.Card
	if err := db.Find(&cards).Error; err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *cardRepository) UpdateByID(id uuid.UUID, updateData dto.UpdateCardRequest) error {
	return r.db.Model(&model.Card{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *cardRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.Card{}).Error
}
