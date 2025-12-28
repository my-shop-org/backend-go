package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   `json:"product_id"`
	VariantID *uint  `json:"variant_id,omitempty"`
	URL       string `gorm:"not null"`
	IsDefault bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
