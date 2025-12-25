package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attribute struct {
	gorm.Model
	ID     uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name   string           `gorm:"unique;not null"`
	Values []AttributeValue `gorm:"foreignKey:AttributeID"`
}
