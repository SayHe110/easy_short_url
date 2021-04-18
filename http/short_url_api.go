package http

import (
	"easy_short_url/pkg/client"
	"easy_short_url/pkg/response"
	"easy_short_url/pkg/short_url"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

var genShortUrlPrefix = "FULLURL:"

func genShortUrl(c *gin.Context) {
	fullUrl := c.PostForm("full_url")

	genedShortUrl := short_url.GenShortUrl(fullUrl)

	redisKey := genShortUrlPrefix + genedShortUrl + ":"

	res, redisErr := client.RedisDb.Get().Do("Set", redisKey, fullUrl)

	if redisErr != nil {
		response.JSON(c, http.StatusOK, "Error", gin.H{
			"Error": redisErr.Error(),
		})

		return
	}

	// 设置过期时间
	_, _ = client.RedisDb.Get().Do("expire", redisKey, 24*60*60)

	response.JSON(c, http.StatusOK, "Success", gin.H{
		"full_url":  fullUrl,
		"short_url": genedShortUrl,
		"res":       res,
	})
}

func redirectFullUrl(c *gin.Context) {
	shortUrl := c.Param("short_url")

	redisKey := genShortUrlPrefix + shortUrl + ":"

	fullUrl, err := redis.String(client.RedisDb.Get().Do("Get", redisKey))

	if err != nil {
		response.JSON(c, 10000, "Error", gin.H{
			"Error": err.Error(),
		})

		return
	}

	if fullUrl == "" {
		response.JSON(c, 10001, "短链不存在，请重新生成", nil)

		return
	}

	c.Redirect(http.StatusMovedPermanently, fullUrl)
}
