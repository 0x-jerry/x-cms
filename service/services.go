package service

type AllServices struct {
	Tag      TagService
	Category CategoryService

	Post         PostService
	PostTag      PostTagService
	PostCategory PostCategoryService

	Article ArticleService
}
