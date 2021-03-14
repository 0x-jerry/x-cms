package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(dbName string) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Category{}, &Note{}, &Post{}, &Tag{})
}

func Db() *gorm.DB {
	return db
}
