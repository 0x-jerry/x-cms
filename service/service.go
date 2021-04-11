package service

import "gorm.io/gorm"

type BasicService struct {
	db *gorm.DB
}

func (s *BasicService) SetDB(db *gorm.DB) {
	s.db = db
}
