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

		db := entity.Db()
		var posts []entity.Post

		db.Order(pagenation.SortBy + " desc").Offset(pagenation.Page * pagenation.Size).Limit(pagenation.Size).Find(&posts)

		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	})
}
