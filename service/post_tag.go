package service

import "github.com/cwxyz007/x-cms/model"

type PostTagService struct {
	BasicService
	TagService TagService
}

func (s *PostTagService) GetByPostIds(postIds []string) ([]model.PostTag, error) {
	var postTags []model.PostTag

	err := s.db.Where("post_id in ?", postIds).Find(&postTags).Error

	return postTags, err
}

func (s *PostTagService) GetByTagIds(tagIds []string) ([]model.PostTag, error) {
	var postTags []model.PostTag

	err := s.db.Where("tag_id in ?", tagIds).Find(&postTags).Error

	return postTags, err
}

func (s *PostTagService) Create(postTag model.PostTag) error {
	return s.db.Create(&postTag).Error
}

func (s *PostTagService) CreateBatch(postTags []model.PostTag) error {
	return s.db.Create(&postTags).Error
}
