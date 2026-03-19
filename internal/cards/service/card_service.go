package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/model"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/repository"
)

type CardService interface {
	CreateCard(req dto.CreateCardRequest) (*model.Card, error)
	GetCard(id uuid.UUID) (*model.Card, error)
	GetByCardNumber(cardNumber string) (*model.Card, error)
	GetAllCards() ([]*model.Card, error)
	GetCards(filter dto.CardFilter) ([]*model.Card, error)
	UpdateCardByID(id uuid.UUID, updateData dto.UpdateCardRequest) error
	DeleteCardByID(id uuid.UUID) error
}

type cardService struct {
	repo repository.CardRepository
}

func NewCardService(r repository.CardRepository) CardService {
	return &cardService{repo: r}
}

func (s *cardService) CreateCard(req dto.CreateCardRequest) (*model.Card, error) {
	existingCard, err := s.repo.GetByCardNumber(req.CardNumber)
	if err == nil && existingCard != nil {
		return nil, errors.New("card with this number already exists")
	}

	keyID, err := uuid.Parse(req.KeyID)
	if err != nil {
		return nil, errors.New("invalid key_id")
	}

	card := &model.Card{
		ID:          uuid.New(),
		CardNumber:  req.CardNumber,
		OwnerName:   req.OwnerName,
		Balance:     req.Balance,
		IsBlocked:   req.IsBlocked,
		BlockReason: req.BlockReason,
		KeyID:       keyID,
		ExpiresAt:   req.ExpiresAt,
	}

	if err := s.repo.Create(card); err != nil {
		return nil, err
	}

	return card, nil
}

func (s *cardService) GetCard(id uuid.UUID) (*model.Card, error) {
	card, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("card not found")
	}
	return card, nil
}

func (s *cardService) GetByCardNumber(cardNumber string) (*model.Card, error) {
	return s.repo.GetByCardNumber(cardNumber)
}

func (s *cardService) GetAllCards() ([]*model.Card, error) {
	cards, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (s *cardService) GetCards(filter dto.CardFilter) ([]*model.Card, error) {
	return s.repo.Find(filter)
}

func (s *cardService) UpdateCardByID(id uuid.UUID, updateData dto.UpdateCardRequest) error {
	if updateData.KeyID != nil {
		parsed, err := uuid.Parse(*updateData.KeyID)
		if err != nil {
			return errors.New("invalid key_id")
		}
		keyID := parsed.String()
		updateData.KeyID = &keyID
	}

	return s.repo.UpdateByID(id, updateData)
}

func (s *cardService) DeleteCardByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
