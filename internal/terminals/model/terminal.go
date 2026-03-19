package model

import (
	"time"

	"github.com/google/uuid"
)

type Terminal struct {
	ID           uuid.UUID  `gorm:"type:text;primaryKey"`
	SerialNumber string     `gorm:"type:text;not null;uniqueIndex"`
	Name         string     `gorm:"type:text;not null"`
	Location     string     `gorm:"type:text;not null"`
	Route        *string    `gorm:"type:text"`
	IsActive     bool       `gorm:"not null;default:true"`
	LastSeenAt   *time.Time `gorm:"type:datetime"`
	CreatedAt    time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`

	// relations
	// Transactions []transactionModel.Transaction `gorm:"foreignKey:TerminalID"`
}
