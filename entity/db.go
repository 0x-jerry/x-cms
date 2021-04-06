package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Initialize(dbName string) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(User{}, Category{}, Note{}, Post{}, Tag{}, PostTag{}, PostCategory{})
	dbInstance = db
}

func Db() *gorm.DB {
	return dbInstance
}

// Basic model
type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
