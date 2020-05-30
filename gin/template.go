package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"lang": "GO语言",
			"tag":  1 / getZero(),
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func getZero() int {
	return 0
}
