package entity

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	ProductID     uint `gorm:"not null"`
	Product       Product
	SKU           string `gorm:"uniqueIndex;not null"`
	Price         float64
	Stock         int
	ProductImages []ProductImage   `gorm:"foreignKey:VariantID"`
	Attributes    []AttributeValue `gorm:"many2many:variant_attribute_values;"`
}
