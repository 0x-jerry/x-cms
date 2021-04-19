package service

import (
	"github.com/cwxyz007/x-cms/model"
)

type CategoryService struct {
	BasicService
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	var categories []model.Category

	err := s.db.Find(&categories).Error

	return categories, err
}

func (s *CategoryService) GetBy(id string) (model.Category, error) {
	var category model.Category

	err := s.db.First(&category, "id == ?", id).Error

	return category, err
}

func (s *CategoryService) GetBatchBy(ids []string) ([]model.Category, error) {
	var categories []model.Category

	err := s.db.Where("id in ?", ids).Find(&categories).Error

	return categories, err
}
