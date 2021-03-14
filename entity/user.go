package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserName string `gorm:"unique;not null;"`
	Password string `gorm:"unique;not null;"`

	NickName string
	Level    int8
}
