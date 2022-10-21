package main

import (
	"github.com/gin-gonic/gin"
)

//响应方式

func main() {
	//json
	r := gin.Default()
	r.GET("json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"xfp": "handsome", "age": 23})
	})
	//结构体
	r.GET("struct", func(ctx *gin.Context) {
		var user struct {
			name  string
			hobby string
			age   int
		}
		user.name = "xfp"
		user.hobby = "swimming"
		user.age = 23
		ctx.JSON(200, user)
	})
	//xml
	r.GET("xml", func(ctx *gin.Context) {
		ctx.XML(200, gin.H{"message": "abc"})
	})
	//yaml
	r.GET("yaml", func(ctx *gin.Context) {
		ctx.YAML(200, gin.H{"name": "徐发鹏"})
	})
	// //protobuf,微服务中的
	// r.GET("protobuf", func(ctx *gin.Context) {
	// 	response := []int64{int64(1), int64(2)}
	// 	label := "mircosoft"
	// 	data := &protoexample.Test
	// })
}
