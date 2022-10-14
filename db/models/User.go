package models

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	Email  string `gorm:"not null" json:"email"`
	Name   string `gorm:"not null" json:"name"`
	UserId string `gorm:"not null" json:"user_id"`
}

type User struct {
	gorm.Model
	UserProfile
	Password string `gorm:"not null"`
}

func (u *User) LoggedIn() bool {
	return u.ID != 0
}

func (u *User) Data() UserProfile {
	return u.UserProfile
}
