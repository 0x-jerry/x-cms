package entity

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	title string

	content string

	tags []Tag

	categories []Category
}
