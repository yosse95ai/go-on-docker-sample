package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string  `josn:"user_name"`
	Password string  `json:"-"`
	Token    *string `json:"token"`
}
