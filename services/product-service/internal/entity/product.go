package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Name          string         `json:"name" gorm:"unique;not null"`
	Description   string         `json:"description"`
	Categories    []*Category    `gorm:"many2many:product_categories;" json:"categories,omitempty"`
	Attributes    []*Attribute   `gorm:"many2many:product_attributes;" json:"attributes,omitempty"`
	BasePrice     float64        `gorm:"not null" json:"base_price"`
	ComparePrice  float64        `json:"compare_price" gorm:"default:0"`
	Currency      string         `gorm:"not null;default:MMK" json:"currency"`
	ProductImages []ProductImage `gorm:"foreignKey:ProductID" json:"product_images,omitempty"`
	Variants      []Variant      `gorm:"foreignKey:ProductID" json:"variants,omitempty"`

	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ComparePrice == 0 {
		p.ComparePrice = p.BasePrice
	}
	return
}
