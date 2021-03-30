package entity

type Tag struct {
	Model

	Name     string `json:"name"`
	Describe string `json:"describe"`
}

func GetTagsByIds(ids []uint) ([]Tag, error) {
	var tags []Tag

	err := Db().Where("id in ?", ids).Find(&tags).Error

	return tags, err
}
