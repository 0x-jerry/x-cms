package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	r := router.Group("/api")

	GetArticles(r)
	GetArticle(r)
	CreateArticle(r)
	UpdateArticle(r)
	DeleteArticle(r)
}
