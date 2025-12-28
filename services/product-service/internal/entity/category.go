package entity

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Name        string      `json:"name" gorm:"unique;not null"`
	Description string      `json:"description,omitempty"`
	ParentID    *uint       `json:"parent_id,omitempty"`
	Children    []*Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Products    []*Product  `json:"products,omitempty" gorm:"many2many:product_categories;"`
}
