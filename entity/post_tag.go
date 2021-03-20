package entity

type PostTag struct {
	PostID string `json:"postId"`
	Post   Post   `json:"post"`

	TagID string `json:"tagId"`
	Tag   Tag    `json:"tag"`
}
