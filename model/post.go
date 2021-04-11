package model

type Post struct {
	Model

	Title   string `json:"title"`
	Summary string `json:"summary"`

	Content string `json:"content"`
}
