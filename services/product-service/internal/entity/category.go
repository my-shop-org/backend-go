package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string      `gorm:"unique;not null"`
	Description string      `json:"description,omitempty"`
	ParentID    *uuid.UUID  `json:"parent_id,omitempty"`
	Children    []*Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Products    []Product   `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}
