package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	ProductID uint      `gorm:"not null"`
	URL       string    `gorm:"not null"`
	IsPrimary bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
