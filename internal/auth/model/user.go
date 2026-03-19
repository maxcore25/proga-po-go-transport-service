package model

import (
	"time"

	"github.com/google/uuid"
)

type KnowledgeLevel string

const (
	KnowledgeLevelBeginner     KnowledgeLevel = "beginner"
	KnowledgeLevelIntermediate KnowledgeLevel = "intermediate"
	KnowledgeLevelAdvanced     KnowledgeLevel = "advanced"
)

type UserRole string

const (
	RoleClient UserRole = "client"
	RoleTutor  UserRole = "tutor"
	RoleAdmin  UserRole = "admin"
)

type User struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName      string         `gorm:"type:varchar(50);not null"`
	LastName       string         `gorm:"type:varchar(50);not null"`
	MiddleName     *string        `gorm:"type:varchar(50)"`
	Email          string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password       string         `gorm:"type:varchar(255);not null"`
	Phone          *string        `gorm:"type:varchar(20)"`
	KnowledgeLevel KnowledgeLevel `gorm:"type:varchar(20);not null"`
	Role           UserRole       `gorm:"type:varchar(20);not null;default:'client'" json:"role"`

	// Tutor-specific fields (optional)
	Rating            *float64 `gorm:"type:decimal(3,2)" json:"rating,omitempty"`
	Portfolio         *string  `gorm:"type:text" json:"portfolio,omitempty"`
	TestimonialsCount *int     `gorm:"type:int" json:"testimonials_count,omitempty"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
