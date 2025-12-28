package entity

import (
	"time"
)

type Product struct {
	Name          string         `json:"name" gorm:"unique;not null"`
	Description   string         `json:"description"`
	Categories    []*Category    `gorm:"many2many:product_categories;" json:"categories,omitempty"`
	Price         float64        `gorm:"not null" json:"price"`
	ProductImages []ProductImage `gorm:"foreignKey:ProductID" json:"product_images,omitempty"`
	Variants      []Variant      `gorm:"foreignKey:ProductID" json:"variants,omitempty"`

	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
