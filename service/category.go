package service

import (
	"github.com/cwxyz007/x-cms/model"
)

type CategoryService struct {
	BasicService
}

func (s *CategoryService) GetCategoriesByIds(ids []string) ([]model.Category, error) {
	var categories []model.Category

	err := s.db.Where("id in ?", ids).Find(&categories).Error

	return categories, err
}
