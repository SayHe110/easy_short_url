package main

import (
	"easy_short_url/http"
	"easy_short_url/pkg/client"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 运行日志数据至日志文件
	file, _ := os.Create("./runtime/log/gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	r := gin.Default()

	// 路由
	http.Router(r)

	// 启用 redis
	client.SetUpRedisClient()


	// 启用 服务
	if err := r.Run(); err != nil {
		panic(err)
	}
}
