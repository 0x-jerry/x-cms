package service

import "github.com/cwxyz007/x-cms/model"

type PostService struct {
	BasicService
}

func (s *PostService) Create(p *model.Post) error {
	return s.db.Create(&p).Error
}

func (s *PostService) Update(p *model.Post) error {
	return s.db.Model(&p).Updates(p).Error
}

func (s *PostService) GetBy(id string, allInformation bool) (*model.Post, error) {
	post := model.Post{
		Model: model.Model{
			ID: id,
		},
	}

	var err error

	if allInformation {
		err = s.db.Find(&post).Error
	} else {
		err = s.db.Omit("content").Find(&post).Error
	}

	return &post, err
}

func (s *PostService) DeleteBy(id string) error {
	post := model.Post{
		Model: model.Model{
			ID: id,
		},
	}

	return s.db.Delete(&post).Error
}

func (s *PostService) GetBatch(page int, size int, sort string) ([]model.Post, error) {
	offset := page * size

	var posts []model.Post

	err := s.db.Omit("content").Order(sort + " desc").Offset(offset).Limit(size).Find(&posts).Error

	if err != nil {
		return posts, err
	}

	return posts, err
}
