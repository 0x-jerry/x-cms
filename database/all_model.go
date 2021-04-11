package database

import "github.com/cwxyz007/x-cms/model"

var AllModels = []interface{}{
	model.User{},
	model.Category{},
	model.Note{},
	model.Post{},
	model.Tag{},
	model.PostTag{},
	model.PostCategory{},
}
