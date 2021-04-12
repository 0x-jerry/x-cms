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
