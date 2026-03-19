package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/model"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/repository"
)

type KeyService interface {
	CreateKey(req dto.CreateKeyRequest) (*model.Key, error)
	GetKey(id uuid.UUID) (*model.Key, error)
	GetByKeyValue(keyValue string) (*model.Key, error)
	GetAllKeys() ([]*model.Key, error)
	GetKeys(filter dto.KeyFilter) ([]*model.Key, error)
	UpdateKeyByID(id uuid.UUID, updateData dto.UpdateKeyRequest) error
	DeleteKeyByID(id uuid.UUID) error
}

type keyService struct {
	repo repository.KeyRepository
}

func NewKeyService(r repository.KeyRepository) KeyService {
	return &keyService{repo: r}
}

func (s *keyService) CreateKey(req dto.CreateKeyRequest) (*model.Key, error) {
	existingKey, err := s.repo.GetByKeyValue(req.KeyValue)
	if err == nil && existingKey != nil {
		return nil, errors.New("key with this value already exists")
	}

	key := &model.Key{
		ID:          uuid.New(),
		Name:        req.Name,
		KeyValue:    req.KeyValue,
		KeyType:     req.KeyType,
		Sector:      req.Sector,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	if err := s.repo.Create(key); err != nil {
		return nil, err
	}

	return key, nil
}

func (s *keyService) GetKey(id uuid.UUID) (*model.Key, error) {
	key, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("key not found")
	}
	return key, nil
}

func (s *keyService) GetByKeyValue(keyValue string) (*model.Key, error) {
	return s.repo.GetByKeyValue(keyValue)
}

func (s *keyService) GetAllKeys() ([]*model.Key, error) {
	keys, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (s *keyService) GetKeys(filter dto.KeyFilter) ([]*model.Key, error) {
	return s.repo.Find(filter)
}

func (s *keyService) UpdateKeyByID(id uuid.UUID, updateData dto.UpdateKeyRequest) error {
	return s.repo.UpdateByID(id, updateData)
}

func (s *keyService) DeleteKeyByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
