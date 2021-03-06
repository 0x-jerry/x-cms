package controller

import (
	"github.com/cwxyz007/x-cms/model"
	"github.com/cwxyz007/x-cms/service"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type ArticleController struct {
	Ctx      iris.Context
	Services *service.AllServices
	Logger   *logrus.Logger
}

type Pagenation struct {
	Page   int    `form:"page" binding:"gte=0"`
	Size   int    `form:"size" binding:"gte=10,lte=100"`
	SortBy string `form:"sortBy" binding:"oneof=created_at updated_at"`
}

func (c *ArticleController) Get() Response {
	pagenation := Pagenation{
		Page:   0,
		Size:   10,
		SortBy: "created_at",
	}

	if err := c.Ctx.ReadQuery(&pagenation); err != nil {
		return BadRequest(err)
	}

	posts, err := c.Services.Post.GetBatch(pagenation.Page, pagenation.Size, pagenation.SortBy)

	if err != nil {
		return BadRequest(err)
	}

	articles, err := c.Services.Article.GetByPosts(posts)

	return ResponseWithError(articles, err)
}

func (c *ArticleController) GetBy(id string) Response {
	post, err := c.Services.Post.GetBy(id, false)

	if err != nil {
		return BadRequest(err)
	}

	articles, err := c.Services.Article.GetByPosts([]model.Post{*post})

	if err != nil {
		return BadRequest(err)
	}

	return ResponseWithError(articles[0], err)
}

type createArticleParam struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Summary string `json:"summary"`

	TagIds      []string `json:"tagIds"`
	CategoryIds []string `json:"categoryIds"`
}

func (c *ArticleController) Post() Response {
	var post createArticleParam

	if err := c.Ctx.ReadJSON(&post); err != nil {
		return BadRequest(err)
	}

	newPost := model.Post{
		Title:   post.Title,
		Content: post.Content,
		Summary: post.Summary,
	}

	if err := c.Services.Post.Create(&newPost); err != nil {
		return BadRequest(err)
	}

	if len(post.TagIds) > 0 {
		var postTags []model.PostTag

		for _, tagId := range post.TagIds {
			postTags = append(postTags, model.PostTag{
				PostID: newPost.ID,
				TagID:  tagId,
			})
		}

		if err := c.Services.PostTag.CreateBatch(postTags); err != nil {
			return BadRequest(err)
		}
	}

	if len(post.CategoryIds) > 0 {
		var postCategories []model.PostCategory

		for _, categoryId := range post.CategoryIds {
			postCategories = append(postCategories, model.PostCategory{
				PostID:     newPost.ID,
				CategoryID: categoryId,
			})
		}

		if err := c.Services.PostCategory.CreateBatch(postCategories); err != nil {
			return BadRequest(err)
		}
	}

	return ResponseWithError(newPost, nil)
}

type updateArticleParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (c *ArticleController) PutBy(id string) Response {
	var params updateArticleParam

	if err := c.Ctx.ReadJSON(&params); err != nil {
		return BadRequest(err)
	}

	updatePost := model.Post{
		Model: model.Model{
			ID: id,
		},

		Title:   params.Title,
		Content: params.Content,
	}

	err := c.Services.Post.Update(&updatePost)

	return ResponseWithError(SuccessResponse, err)
}
