package entity

type User struct {
	Model

	UserName string `json:"username" gorm:"unique;not null;"`
	Password string `json:"password" gorm:"unique;not null;"`

	NickName string `json:"nickname"`
	Level    int8   `json:"level"`
}
