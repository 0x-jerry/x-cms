package entity

type Post struct {
	Model

	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostTag struct {
	PostID string `json:"postId"`
	Post   Post   `json:"post"`

	TagID string `json:"tagId"`
	Tag   Tag    `json:"tag"`
}

type PostCategory struct {
	PostID string `json:"postId"`
	Post   Post   `json:"post"`

	CategoryID string   `json:"categoryId"`
	Category   Category `json:"category"`
}
