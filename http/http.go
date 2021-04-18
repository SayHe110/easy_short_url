package http

import "github.com/gin-gonic/gin"

func Router(e *gin.Engine) {
	api := e.Group("/api")
	{
		api.POST("/gen_short_url", genShortUrl)
	}

	e.GET("/:short_url", redirectFullUrl)
}
