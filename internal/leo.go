package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个新的 Gin 路由引擎
	r := gin.Default()

	// 定义一个简单的路由，处理 GET 请求
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 启动服务器，默认端口为 8080
	_ = r.Run()

}
