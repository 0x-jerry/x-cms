package model

type Note struct {
	Model

	Title   string `json:"title"`
	Content string `json:"content"`
}
