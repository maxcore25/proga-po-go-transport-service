package model

import (
	"time"

	"github.com/google/uuid"

	cardModel "github.com/maxcore25/proga-po-go-transport-service/internal/cards/model"
	terminalModel "github.com/maxcore25/proga-po-go-transport-service/internal/terminals/model"
)

type Transaction struct {
	ID            uuid.UUID `gorm:"type:text;primaryKey"`
	CardID        uuid.UUID `gorm:"type:text;not null;index"`
	TerminalID    uuid.UUID `gorm:"type:text;not null;index"`
	Amount        int       `gorm:"not null"`
	BalanceBefore int       `gorm:"not null"`
	BalanceAfter  int       `gorm:"not null"`
	Status        string    `gorm:"type:text;not null;default:approved"`
	DeclineReason *string   `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`

	// relations
	Card     cardModel.Card         `gorm:"foreignKey:CardID;constraint:OnDelete:RESTRICT"`
	Terminal terminalModel.Terminal `gorm:"foreignKey:TerminalID;constraint:OnDelete:RESTRICT"`
}
