package mapper

import (
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/model"
)

func NewCardResponse(c *model.Card) *dto.CardResponse {
	return &dto.CardResponse{
		ID:          c.ID,
		CardNumber:  c.CardNumber,
		OwnerName:   c.OwnerName,
		Balance:     c.Balance,
		IsBlocked:   c.IsBlocked,
		BlockReason: c.BlockReason,
		KeyID:       c.KeyID,
		ExpiresAt:   c.ExpiresAt,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
