package entity

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ID         uint             `gorm:"primaryKey"`
	ProductID  uint             `gorm:"not null"`
	Product    Product
	SKU        string           `gorm:"uniqueIndex;not null"`
	Price      float64
	Stock      int
	Attributes []AttributeValue `gorm:"many2many:variant_attribute_values;"`
}
