package main

import (
	"time"

	"github.com/cwxyz007/x-cms/api"
	"github.com/cwxyz007/x-cms/core"
	"github.com/cwxyz007/x-cms/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := core.GetConfig()

	entity.Initialize("test.db")

	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(logger())

	api.RegisterRoutes(router)

	router.Run(conf.Port)
}

// logger instances a logger middleware for Gin.
func logger() gin.HandlerFunc {
	log := core.GetLogger()

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		// Use debug level to keep production logs clean.
		log.Debugf("http: %s (%d) [%v] %s",
			method,
			statusCode,
			latency,
			path,
		)
	}
}
