package entity

type Post struct {
	Model

	Title   string `json:"title"`
	Summary string `json:"summary"`

	Content string `json:"content"`
}

type Article struct {
	Post
	Tags       []Tag      `json:"tags"`
	Categories []Category `json:"categories"`
}

func CreatePost(p *Post) error {
	return Db().Create(&p).Error
}

func UpdatePost(p *Post) error {
	return Db().Model(&p).Updates(p).Error
}

func GetPost(id uint, allInformation bool) (*Article, error) {
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

	articles, err := GetArticles([]Post{post})

	if err != nil {
		return nil, err
	}

	return &articles[0], err
}

func DeletePost(id uint) error {
	post := Post{
		Model: Model{
			ID: id,
		},
	}

	return Db().Delete(&post).Error
}

func GetPosts(page int, size int, sort string) ([]Article, error) {
	offset := page * size

	var posts []Post

	articles := []Article{}

	err := Db().Omit("content").Order(sort + " desc").Offset(offset).Limit(size).Find(&posts).Error

	if err != nil {
		return articles, err
	}

	articles, err = GetArticles(posts)

	return articles, err
}

func GetArticles(posts []Post) ([]Article, error) {
	var articles []Article

	var (
		postIds         []uint
		postTagIds      []uint
		postCategoryIds []uint
	)

	for _, v := range posts {
		postIds = append(postIds, v.ID)
	}

	// 获取 tag, category 和 post 的关系
	postTags, err := GetTagsByPostIds(postIds)
	postCategories, err := GetCategoriesByPostIds(postIds)

	if err != nil {
		return articles, err
	}

	for _, v := range postTags {
		postTagIds = append(postTagIds, v.TagID)
	}

	for _, v := range postCategories {
		postCategoryIds = append(postTagIds, v.CategoryID)
	}

	// 获取有关联的所有 tag 和 category
	tags, err := GetTagsByIds(postTagIds)
	categories, err := GetCategoriesByIds(postCategoryIds)

	if err != nil {
		return articles, err
	}

	for _, v := range posts {
		// 找到每一个 post 对应的 tags 和 categories
		pTags := findTagsByPostId(postTags, tags, v)
		pCategories := findCategoriesByPostId(postCategories, categories, v)

		articles = append(articles, Article{
			Post:       v,
			Tags:       pTags,
			Categories: pCategories,
		})
	}

	return articles, err
}

// 根据 post 和 category 的关系，找到这个 post 所有的 categories
func findCategoriesByPostId(postTags []PostCategory, categories []Category, post Post) []Category {
	pCategories := []Category{}

	for _, v := range postTags {
		if v.PostID == post.ID {
			pCategories = append(pCategories, findCategoryById(categories, v.CategoryID))
		}
	}

	return pCategories
}

func findCategoryById(categories []Category, id uint) Category {
	var category Category

	for _, v := range categories {
		if v.ID == id {
			category = v
			break
		}
	}

	return category
}

// 根据 post 和 tag 的关系，找到这个 post 所有的 tags
func findTagsByPostId(postTags []PostTag, tags []Tag, post Post) []Tag {
	pTags := []Tag{}

	for _, v := range postTags {
		if v.PostID == post.ID {
			pTags = append(pTags, findTagById(tags, v.TagID))
		}
	}

	return pTags
}

func findTagById(tags []Tag, id uint) Tag {
	var tag Tag

	for _, v := range tags {
		if v.ID == id {
			tag = v
			break
		}
	}

	return tag
}
