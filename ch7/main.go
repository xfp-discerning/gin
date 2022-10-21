package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		ctx.Set("request", "执行")
		ctx.Next()                                  //执行
		fmt.Println("中间件执行完毕", ctx.Writer.Status()) //200
		t2 := time.Since(t)
		fmt.Println("用时：", t2)
	}
}

func main() {
	r := gin.Default()
	//注册中间件
	r.Use(MiddleWare())
	//{}是为了代码规范
	{
		r.GET("/middleware", func(ctx *gin.Context) {
			req, _ := ctx.Get("request")
			fmt.Println(req)
			ctx.JSON(200, gin.H{"requst": req})
		})

		//跟路由后面，是定义的局部中间件
		r.GET("/middleware2",MiddleWare(), func(ctx *gin.Context) {
			req, _ := ctx.Get("request")
			fmt.Println(req)
			ctx.JSON(200, gin.H{"requst": req})
		})

	}
	r.Run(":8000")
}
