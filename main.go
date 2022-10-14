package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则
	//gin.Context封装了request和response
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "xfp")
	})
	//监听端口
	r.Run("localhost:8000")
}
