package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
c.ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。
这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法。
只有某些格式需要此功能，如 JSON, XML, MsgPack, ProtoBuf。 对于其他格式, 如
Query, Form, FormPost, FormMultipart 可以多次调用 c.ShouldBind()
而不会造成任任何性能损失 (详见 #1341)。
*/
func main() {
	engine := gin.Default()
	engine.POST("/", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}
		// c.ShouldBind 使用了 c.Request.Body，不可重用。
		if errA := c.ShouldBind(&objA); errA == nil {
			c.String(http.StatusOK, `the body should be formA`)
			// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		} else if errB := c.ShouldBind(&objB); errB == nil {
			c.String(http.StatusOK, `the body should be formB`)
		} else {
		}
	})
	engine.Run()
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}
