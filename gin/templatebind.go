package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Type - Must bind
Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML
Behavior - 这些方法属于 MustBindWith 的具体调用。 如果发生绑定错误，则请求终止，
并触发 c.AbortWithError(400, err).SetType(ErrorTypeBind)。
响应状态码被设置为 400 并且 Content-Type 被设置为 text/plain; charset=utf-8。
如果您在此之后尝试设置响应状态码，Gin会输出日志 [GIN-debug] [WARNING] Headers were already written.
Wanted to override status code 400 with 422。 如果您希望更好地控制绑定，考虑使用 ShouldBind 等效方法。

Type - Should bind
Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML
Behavior - 这些方法属于 ShouldBindWith 的具体调用。 如果发生绑定错误，Gin 会返回错误并由开发者处理错误和请求。

*/
// 绑定 JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 绑定 JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 绑定 XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 绑定 HTML 表单 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 根据 Content-Type Header 推断使用哪个绑定器。
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
