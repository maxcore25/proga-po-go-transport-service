package mapper

import (
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/model"
)

func NewTransactionResponse(t *model.Transaction) *dto.TransactionResponse {
	return &dto.TransactionResponse{
		ID:            t.ID,
		CardID:        t.CardID,
		TerminalID:    t.TerminalID,
		Amount:        t.Amount,
		BalanceBefore: t.BalanceBefore,
		BalanceAfter:  t.BalanceAfter,
		Status:        t.Status,
		DeclineReason: t.DeclineReason,
		CreatedAt:     t.CreatedAt,
	}
}
