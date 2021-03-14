package entity

import "gorm.io/gorm"

type Tag struct {
	gorm.Model

	name     string
	describe string
}
