package service

import "github.com/cwxyz007/x-cms/model"

type ArticleService struct {
	BasicService

	TagService      TagService
	CategoryService CategoryService

	PostTagService      PostTagService
	PostCategoryService PostCategoryService
}

// 获取 tags 以及 categories，然后组装成 articles
func (s *ArticleService) GetByPosts(posts []model.Post) ([]model.Article, error) {
	articles := []model.Article{}

	var (
		postIds         []string
		postTagIds      []string
		postCategoryIds []string
	)

	for _, v := range posts {
		postIds = append(postIds, v.ID)
	}

	// 获取 tag, category 和 post 的关系
	postTags, err := s.PostTagService.GetByPostIds(postIds)
	if err != nil {
		return articles, err
	}

	postCategories, err := s.PostCategoryService.GetByPostIds(postIds)
	if err != nil {
		return articles, err
	}

	for _, v := range postTags {
		postTagIds = append(postTagIds, v.TagID)
	}

	for _, v := range postCategories {
		postCategoryIds = append(postCategoryIds, v.CategoryID)
	}

	// 获取有关联的所有 tag 和 category
	tags, err := s.TagService.GetBy(postTagIds)
	if err != nil {
		return articles, err
	}

	categories, err := s.CategoryService.GetBy(postCategoryIds)
	if err != nil {
		return articles, err
	}

	for _, v := range posts {
		// 找到每一个 post 对应的 tags 和 categories
		pTags := findTagsByPostId(postTags, tags, v)
		pCategories := findCategoriesByPostId(postCategories, categories, v)

		articles = append(articles, model.Article{
			Post:       v,
			Tags:       pTags,
			Categories: pCategories,
		})
	}

	return articles, err
}

// 根据 post 和 category 的关系，找到这个 post 所有的 categories
func findCategoriesByPostId(postTags []model.PostCategory, categories []model.Category, post model.Post) []model.Category {
	pCategories := []model.Category{}

	for _, v := range postTags {
		if v.PostID == post.ID {
			pCategories = append(pCategories, findCategoryById(categories, v.CategoryID))
		}
	}

	return pCategories
}

func findCategoryById(categories []model.Category, id string) model.Category {
	var category model.Category

	for _, v := range categories {
		if v.ID == id {
			category = v
			break
		}
	}

	return category
}

// 根据 post 和 tag 的关系，找到这个 post 所有的 tags
func findTagsByPostId(postTags []model.PostTag, tags []model.Tag, post model.Post) []model.Tag {
	pTags := []model.Tag{}

	for _, v := range postTags {
		if v.PostID == post.ID {
			pTags = append(pTags, findTagById(tags, v.TagID))
		}
	}

	return pTags
}

func findTagById(tags []model.Tag, id string) model.Tag {
	var tag model.Tag

	for _, v := range tags {
		if v.ID == id {
			tag = v
			break
		}
	}

	return tag
}
