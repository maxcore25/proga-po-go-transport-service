package model

import (
	"time"

	"github.com/google/uuid"
	keyModel "github.com/maxcore25/proga-po-go-transport-service/internal/keys/model"
)

type Card struct {
	ID          uuid.UUID  `gorm:"type:text;primaryKey"`
	CardNumber  string     `gorm:"type:text;not null;uniqueIndex"`
	OwnerName   string     `gorm:"type:text;not null"`
	Balance     int        `gorm:"not null;default:0"`
	IsBlocked   bool       `gorm:"not null;default:false"`
	BlockReason *string    `gorm:"type:text"`
	KeyID       uuid.UUID  `gorm:"type:text;not null;index"`
	ExpiresAt   *time.Time `gorm:"type:datetime"`
	CreatedAt   time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`

	// relations
	Key keyModel.Key `gorm:"foreignKey:KeyID;constraint:OnDelete:RESTRICT"`
}
