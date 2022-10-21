package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*")
	// r.LoadHTMLFiles("template/index.tmpl")
	r.GET("/index", func(ctx *gin.Context) {
		//根据文件名渲染
		ctx.HTML(200, "index.tmpl", gin.H{"title": "my ass"})
	})
	r.Run(":8000")
}
