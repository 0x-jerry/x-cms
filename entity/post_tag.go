package entity

type PostTag struct {
	Model

	PostID uint `json:"postId"`
	Post   Post `json:"post"`

	TagID uint `json:"tagId"`
	Tag   Tag  `json:"tag"`
}

func GetTagsByPostIds(postIds []uint) (*[]PostTag, error) {
	var postTags []PostTag

	err := Db().Where("post_id in ?", postIds).Find(&postTags).Error

	return &postTags, err
}
