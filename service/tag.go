package service

import "github.com/cwxyz007/x-cms/model"

type TagService struct {
	BasicService
}

func (s *TagService) GetBy(id string) (model.Tag, error) {
	var tag model.Tag

	err := s.db.Where("id == ?", id).Find(&tag).Error

	return tag, err
}

func (s *TagService) GetBatchBy(ids []string) ([]model.Tag, error) {
	var tags []model.Tag

	err := s.db.Where("id in ?", ids).Find(&tags).Error

	return tags, err
}

func (s *TagService) GetAll() ([]model.Tag, error) {
	var tags []model.Tag

	err := s.db.Find(&tags).Error

	return tags, err
}
