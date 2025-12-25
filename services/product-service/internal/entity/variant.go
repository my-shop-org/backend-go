package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductID  uuid.UUID `gorm:"type:uuid;not null"`
	Product    Product
	SKU        string `gorm:"uniqueIndex;not null"`
	Price      float64
	Stock      int
	Attributes []VariantAttributeValue
}
