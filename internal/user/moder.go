package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index;unique;not null"`
	Password string `gorm:"not null"`
	Name     string
}
