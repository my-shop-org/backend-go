package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VariantAttributeValue struct {
	gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	VariantID        uuid.UUID `gorm:"type:uuid;not null"`
	Variant          Variant
	AttributeValueID uuid.UUID `gorm:"type:uuid;not null"`
	AttributeValue   AttributeValue
}
