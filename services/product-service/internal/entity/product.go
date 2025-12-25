package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	CategoryID  uuid.UUID `gorm:"not null"`
	Category    Category
	Price       float64        `gorm:"not null"`
	CreatedBy   uuid.UUID      `gorm:"not null"`
	Images      []ProductImage `gorm:"foreignKey:ProductID"`
	Variants    []Variant      `gorm:"foreignKey:ProductID"`
}
