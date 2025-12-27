package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"not null"`
	Description string
	CategoryID  uint           `gorm:"not null"`
	Category    Category
	Price       float64        `gorm:"not null"`
	CreatedBy   uint           `gorm:"not null"`
	Images      []ProductImage `gorm:"foreignKey:ProductID"`
	Variants    []Variant      `gorm:"foreignKey:ProductID"`
}
