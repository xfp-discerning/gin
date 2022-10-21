package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("loginJson", func(ctx *gin.Context) {
		//声明接收变量
		var json login
		//将request中的body自动解析到json的结构中
		if err := ctx.ShouldBindJSON(&json); err != nil {
			//gin.H封装了生成json的工具
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//校验密码
		if json.User != "root" || json.Password != "1234" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.POST("form", func(ctx *gin.Context) {
		var form login
		if err := ctx.Bind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "root" || form.Password != "1234" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.GET("/:user/:password", func(ctx *gin.Context) {
		var url login
		if err := ctx.ShouldBindUri(&url); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if url.User != "root" || url.Password != "1234" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}
