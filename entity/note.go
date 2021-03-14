package entity

import "gorm.io/gorm"

type Note struct {
	gorm.Model

	title   string
	content string
}
