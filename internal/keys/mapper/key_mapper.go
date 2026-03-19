package mapper

import (
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/model"
)

func NewKeyResponse(k *model.Key) *dto.KeyResponse {
	return &dto.KeyResponse{
		ID:          k.ID,
		Name:        k.Name,
		KeyValue:    k.KeyValue,
		KeyType:     k.KeyType,
		Sector:      k.Sector,
		Description: k.Description,
		IsActive:    k.IsActive,
		CreatedAt:   k.CreatedAt,
		UpdatedAt:   k.UpdatedAt,
	}
}
