package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/model"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/repository"
)

type TerminalService interface {
	CreateTerminal(req dto.CreateTerminalRequest) (*model.Terminal, error)
	GetTerminal(id uuid.UUID) (*model.Terminal, error)
	GetBySerialNumber(serialNumber string) (*model.Terminal, error)
	GetAllTerminals() ([]*model.Terminal, error)
	GetTerminals(filter dto.TerminalFilter) ([]*model.Terminal, error)
	UpdateTerminalByID(id uuid.UUID, updateData dto.UpdateTerminalRequest) error
	DeleteTerminalByID(id uuid.UUID) error
}

type terminalService struct {
	repo repository.TerminalRepository
}

func NewTerminalService(r repository.TerminalRepository) TerminalService {
	return &terminalService{repo: r}
}

func (s *terminalService) CreateTerminal(req dto.CreateTerminalRequest) (*model.Terminal, error) {
	existingTerminal, err := s.repo.GetBySerialNumber(req.SerialNumber)
	if err == nil && existingTerminal != nil {
		return nil, errors.New("terminal with this serial number already exists")
	}

	terminal := &model.Terminal{
		ID:           uuid.New(),
		SerialNumber: req.SerialNumber,
		Name:         req.Name,
		Location:     req.Location,
		Route:        req.Route,
		IsActive:     req.IsActive,
		LastSeenAt:   req.LastSeenAt,
	}

	if err := s.repo.Create(terminal); err != nil {
		return nil, err
	}

	return terminal, nil
}

func (s *terminalService) GetTerminal(id uuid.UUID) (*model.Terminal, error) {
	terminal, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("terminal not found")
	}
	return terminal, nil
}

func (s *terminalService) GetBySerialNumber(serialNumber string) (*model.Terminal, error) {
	return s.repo.GetBySerialNumber(serialNumber)
}

func (s *terminalService) GetAllTerminals() ([]*model.Terminal, error) {
	terminals, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return terminals, nil
}

func (s *terminalService) GetTerminals(filter dto.TerminalFilter) ([]*model.Terminal, error) {
	return s.repo.Find(filter)
}

func (s *terminalService) UpdateTerminalByID(id uuid.UUID, updateData dto.UpdateTerminalRequest) error {
	return s.repo.UpdateByID(id, updateData)
}

func (s *terminalService) DeleteTerminalByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
