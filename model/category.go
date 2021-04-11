package model

type Category struct {
	Model
	Name     string `json:"name"`
	Describe string `json:"describe"`
}
