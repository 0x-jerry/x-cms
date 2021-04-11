package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Basic model
type Model struct {
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (m *Model) BeforeCreate(db *gorm.DB) (err error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	m.ID = uuid.String()

	return
}
