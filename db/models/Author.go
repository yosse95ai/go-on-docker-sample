package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Books []Book `gorm:"many2many:author_books" json:"books"`
}
