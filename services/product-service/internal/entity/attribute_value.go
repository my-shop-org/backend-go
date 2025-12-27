package entity

import (
	"time"

	"gorm.io/gorm"
)

type AttributeValue struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	AttributeID uint      `gorm:"not null"`
	Attribute   Attribute
	Value       string    `gorm:"not null"` // e.g., S, M, L or Red, Blue
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
