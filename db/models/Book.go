package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string    `gorm:"not null" json:"title"`
	PublisherID uint      `gorm:"not null" json:"publisher_id"`
	Publisher   Publisher `json:"publisher"`
	Authors     []Author  `gorm:"many2many:author_books" json:"authors"`
}
