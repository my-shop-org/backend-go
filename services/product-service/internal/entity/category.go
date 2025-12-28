package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string      `gorm:"unique;not null"`
	Description string      `json:"description,omitempty"`
	ParentID    *uint       `json:"parent_id,omitempty"`
	Children    []*Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Products    []*Product  `json:"products,omitempty" gorm:"many2many:product_categories;"`
}
