package entity

type Category struct {
	Model
	Name     string `json:"name"`
	Describe string `json:"describe"`
}

func GetCategoriesByIds(ids []uint) ([]Category, error) {
	var categories []Category

	err := Db().Where("id in ?", ids).Find(&categories).Error

	return categories, err
}
