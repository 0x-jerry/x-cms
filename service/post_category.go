package service

import "github.com/cwxyz007/x-cms/model"

type PostCategoryService struct {
	BasicService
	CategoryService CategoryService
}

func (s *PostCategoryService) GetByPostIds(postIds []string) ([]model.PostCategory, error) {
	var postCategories []model.PostCategory

	err := s.db.Where("post_id in ?", postIds).Find(&postCategories).Error

	return postCategories, err
}

func (s *PostCategoryService) Create(postCategory model.PostCategory) error {
	return s.db.Create(&postCategory).Error
}

func (s *PostCategoryService) CreateBatch(postCategories []model.PostCategory) error {
	return s.db.Create(&postCategories).Error
}
