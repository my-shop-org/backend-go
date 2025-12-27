package entity

import (
	"gorm.io/gorm"
)

type VariantAttributeValue struct {
	gorm.Model
	ID               uint `gorm:"primaryKey"`
	VariantID        uint `gorm:"not null"`
	Variant          Variant
	AttributeValueID uint `gorm:"not null"`
	AttributeValue   AttributeValue
}
