package entity

type PostCategory struct {
	PostID uint `json:"postId"`
	Post   Post `json:"post"`

	CategoryID uint     `json:"categoryId"`
	Category   Category `json:"category"`
}

func GetCategoriesByPostIds(postIds []uint) (*[]PostCategory, error) {
	var postCategories []PostCategory

	err := Db().Where("post_id in ?", postIds).Find(&postCategories).Error

	return &postCategories, err
}
