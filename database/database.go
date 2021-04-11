package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(dbName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(AllModels...)

	return db
}
