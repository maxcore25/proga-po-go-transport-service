package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:text;primaryKey"`
	Username     string    `gorm:"type:text;not null;uniqueIndex"`
	PasswordHash string    `gorm:"type:text;not null"`
	FullName     string    `gorm:"type:text;not null"`
	IsAdmin      bool      `gorm:"not null;default:false"`
	IsActive     bool      `gorm:"not null;default:true"`
	CreatedAt    time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}
