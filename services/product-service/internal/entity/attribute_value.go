package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttributeValue struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	AttributeID uuid.UUID `gorm:"type:uuid;not null"`
	Attribute   Attribute
	Value       string `gorm:"not null"` // e.g., S, M, L or Red, Blue
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
