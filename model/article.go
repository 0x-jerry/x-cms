package model

type Article struct {
	Post
	Tags       []Tag      `json:"tags"`
	Categories []Category `json:"categories"`
}
