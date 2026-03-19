package model

import (
	"time"

	"github.com/google/uuid"
)

type Key struct {
	ID          uuid.UUID `gorm:"type:text;primaryKey"`
	Name        string    `gorm:"type:text;not null"`
	KeyValue    string    `gorm:"type:text;not null"` // hex
	KeyType     string    `gorm:"type:text;not null;default:A"`
	Sector      int       `gorm:"not null;default:0"`
	Description *string   `gorm:"type:text"`
	IsActive    bool      `gorm:"not null;default:true"`
	CreatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`

	// relations
	// Cards []cardModel.Card `gorm:"foreignKey:KeyID"`
}
