package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPost(router *gin.RouterGroup) {
	router.GET("posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": "cool"})
	})
}
