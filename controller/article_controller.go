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

func (c *ArticleController) Get() Any {
	pagenation := Pagenation{
		Page:   0,
		Size:   10,
		SortBy: "created_at",
	}

	if err := c.Ctx.ReadQuery(&pagenation); err != nil {
		return BadRequest(err)
	}

	posts, err := c.Services.Post.GetPosts(pagenation.Page, pagenation.Size, pagenation.SortBy)

	if err != nil {
		return BadRequest(err)
	}

	articles, err := c.Services.Article.Get(posts)

	return ResponseWithError(articles, err)
}

func (c *ArticleController) GetBy(id string) Any {
	post, err := c.Services.Post.GetPost(id, false)

	if err != nil {
		return BadRequest(err)
	}

	articles, err := c.Services.Article.Get([]model.Post{*post})

	if err != nil {
		return BadRequest(err)
	}

	return ResponseWithError(articles[0], err)
}

type createArticleParam struct {
	Title   string `json:"title" binding:"required"`
	Summary string `json:"summary"`
	Content string `json:"content" binding:"required"`
}

func (c *ArticleController) Post() Any {
	var post createArticleParam

	if err := c.Ctx.ReadJSON(&post); err != nil {
		return BadRequest(err)
	}

	newPost := model.Post{
		Title:   post.Title,
		Content: post.Content,
		Summary: post.Summary,
	}

	err := c.Services.Post.CreatePost(&newPost)

	return ResponseWithError(newPost, err)
}

type updateArticleParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (c *ArticleController) PutBy(id string) Any {
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

	err := c.Services.Post.UpdatePost(&updatePost)

	return ResponseWithError(SuccessResponse, err)
}
