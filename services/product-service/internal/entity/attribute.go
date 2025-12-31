package entity

import "time"

type Attribute struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Name   string           `gorm:"unique;not null" json:"name"`
	Values []AttributeValue `gorm:"foreignKey:AttributeID" json:"values,omitempty"`
}
