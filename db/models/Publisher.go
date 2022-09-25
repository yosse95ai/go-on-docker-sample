package models

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Books []Book `json:"books"`
}
