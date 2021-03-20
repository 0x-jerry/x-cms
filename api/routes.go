package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	r := router.Group("/api")

	GetPosts(r)
	GetPost(r)
	CreatePost(r)
	UpdatePost(r)
	DeletePost(r)
}
