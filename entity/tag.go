package entity

type Tag struct {
	Model

	Name     string `json:"name"`
	Describe string `json:"describe"`
}
