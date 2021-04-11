package service

import "github.com/cwxyz007/x-cms/model"

type TagService struct {
	BasicService
}

func (s *TagService) GetTagsByIds(ids []string) ([]model.Tag, error) {
	var tags []model.Tag

	err := s.db.Where("id in ?", ids).Find(&tags).Error

	return tags, err
}
