package api

import (
	"net/http"

	"github.com/cwxyz007/x-cms/entity"
	"github.com/gin-gonic/gin"
)

type Pagenation struct {
	Page   int    `form:"page" binding:"gte=0"`
	Size   int    `form:"size" binding:"gte=10,lte=100"`
	SortBy string `form:"sortBy" binding:"oneof=created_at updated_at"`
}

// 获取文章列表
func GetPosts(router *gin.RouterGroup) {
	router.GET("/posts", func(c *gin.Context) {
		pagenation := Pagenation{
			Page:   0,
			Size:   10,
			SortBy: "created_at",
		}

		if err := c.Bind(&pagenation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		posts, err := entity.GetPosts(pagenation.Page, pagenation.Size, pagenation.SortBy)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		articles, err := entity.GetArticles(posts)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"posts": articles,
		})
	})
}

type createPostParam struct {
	Title   string `form:"title" binding:"required"`
	Summary string `form:"summary"`
	Content string `form:"content" binding:"required"`
}

// 创建新的文章
func CreatePost(router *gin.RouterGroup) {
	router.POST("/post", func(c *gin.Context) {
		var post createPostParam
		if err := c.Bind(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		newPost := entity.Post{
			Title:   post.Title,
			Content: post.Content,
			Summary: post.Summary,
		}

		err := entity.CreatePost(&newPost)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"post": newPost,
		})
	})
}

type updatePostParam struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

// 更新文章内容
func UpdatePost(router *gin.RouterGroup) {
	router.PUT("/post/:id", func(c *gin.Context) {
		var params updatePostParam

		if err := c.Bind(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		postId := c.Param("id")

		id, err := ConvertToID(postId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		updatePost := entity.Post{
			Model: entity.Model{
				ID: id,
			},

			Title:   params.Title,
			Content: params.Content,
		}

		err = entity.UpdatePost(&updatePost)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}

// 删除文章
func DeletePost(router *gin.RouterGroup) {
	router.DELETE("/post/:id", func(c *gin.Context) {
		postId := c.Param("id")

		id, err := ConvertToID(postId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		err = entity.DeletePost(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}

// 获取文章详情
func GetPost(router *gin.RouterGroup) {
	router.GET("post/:id", func(c *gin.Context) {
		postId := c.Param("id")

		id, err := ConvertToID(postId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		post, err := entity.GetPost(id, true /* getAllInformation */)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		article, err := entity.GetArticles([]entity.Post{*post})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"post": article,
		})
	})
}
