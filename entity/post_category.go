package entity

type PostCategory struct {
	PostID string `json:"postId"`
	Post   Post   `json:"post"`

	CategoryID string   `json:"categoryId"`
	Category   Category `json:"category"`
}
