package mapper

import (
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/model"
)

func NewTerminalResponse(t *model.Terminal) *dto.TerminalResponse {
	return &dto.TerminalResponse{
		ID:           t.ID,
		SerialNumber: t.SerialNumber,
		Name:         t.Name,
		Location:     t.Location,
		Route:        t.Route,
		IsActive:     t.IsActive,
		LastSeenAt:   t.LastSeenAt,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
	}
}
