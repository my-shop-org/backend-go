package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	URL       string    `gorm:"not null"`
	IsPrimary bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
