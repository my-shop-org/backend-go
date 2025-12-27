package entity

import (
	"gorm.io/gorm"
)

type Attribute struct {
	gorm.Model
	ID     uint             `gorm:"primaryKey"`
	Name   string           `gorm:"unique;not null"`
	Values []AttributeValue `gorm:"foreignKey:AttributeID"`
}
