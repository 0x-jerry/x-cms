package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	r := router.Group("/api")

	GetPost(r)
}
