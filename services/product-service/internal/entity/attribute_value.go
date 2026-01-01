package entity

import (
	"time"
)

type AttributeValue struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	AttributeID uint      `gorm:"not null" json:"attribute_id"`
	Attribute   Attribute `json:"-"`
	Value       string    `gorm:"not null" json:"value"` // e.g., S, M, L or Red, Blue
}
