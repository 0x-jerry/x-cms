package entity

type Post struct {
	Model

	Title   string `json:"title"`
	Summary string `json:"summary"`

	Content string `json:"content"`
}

func CreatePost(p *Post) error {
	return Db().Create(&p).Error
}

func UpdatePost(p *Post) error {
	return Db().Model(&p).Updates(p).Error
}

func GetPost(id uint, allInformation bool) (*Post, error) {
	post := Post{
		Model: Model{
			ID: id,
		},
	}

	var err error

	if allInformation {
		err = Db().Find(&post).Error
	} else {
		err = Db().Omit("content").Find(&post).Error
	}

	return &post, err
}

func DeletePost(id uint) error {
	post := Post{
		Model: Model{
			ID: id,
		},
	}

	return Db().Delete(&post).Error
}

func GetPosts(page int, size int, sort string) ([]Post, error) {
	offset := page * size

	var posts []Post

	err := Db().Omit("content").Order(sort + " desc").Offset(offset).Limit(size).Find(&posts).Error

	if err != nil {
		return posts, err
	}

	return posts, err
}
