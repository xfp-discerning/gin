package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建默认路由
	//默认使用logger() recovery()中间件
	r := gin.Default()
	//router group1 is to deal with GET request
	v1 := r.Group("/v1")
	{
		v1.GET("login", login)
		v1.GET("submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("login", login)
		v2.POST("submit", submit)
	}
	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "xfp")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "xsy")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
