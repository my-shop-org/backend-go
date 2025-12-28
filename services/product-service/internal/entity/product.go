package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null"`
	Description   string
	Categories    []*Category    `json:"categories,omitempty" gorm:"many2many:product_categories;"`
	Price         float64        `gorm:"not null"`
	ProductImages []ProductImage `gorm:"foreignKey:ProductID"`
	Variants      []Variant      `gorm:"foreignKey:ProductID"`
}
