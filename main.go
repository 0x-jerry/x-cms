package main

import (
	"x-cms/core"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	core.Initialize()

	mode := viper.GetString("GIN_MODE")
	gin.SetMode(mode)

	r := gin.Default()

	port := viper.GetString("port")
	r.Run(port)
}
