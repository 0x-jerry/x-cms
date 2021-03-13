package entities

import "gorm.io/gorm"

// admin user model
type User struct {
	gorm.Model

	UserName string
	Password string

	NickName string
	Level    int8
}
