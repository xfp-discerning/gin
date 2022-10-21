package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// //1.创建路由
	// r := gin.Default()
	// //2.绑定路由规则
	// //gin.Context封装了request和response
	// r.GET("/user/:name/*action", func(ctx *gin.Context) {//冒号很有必要
	// 	name := ctx.Param("name")
	// 	// action := ctx.Param("action")
	// 	// ctx.String(http.StatusOK, name+"is"+action)
	// 	ctx.String(http.StatusOK, name)
	// })
	// //监听端口
	// r.Run("localhost:8000")

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(ctx *gin.Context) {
		//表单取文件
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.String(200, fmt.Sprintln(err.Error()))
		}
		//获取所有图片
		files := form.File["file"]
		for _, file := range files {
			if err := ctx.SaveUploadedFile(file, file.Filename); err != nil {
				ctx.String(200, fmt.Sprintf("upload failed, err is %s", err.Error()))
				return
			}
		}
		ctx.String(200, fmt.Sprintf("upload success, number is %d", len(files)))
		// file, _ := ctx.FormFile("file")
		// log.Println(file.Filename)
		// ctx.SaveUploadedFile(file, file.Filename)
		// ctx.String(http.StatusOK, fmt.Sprintf("'%s' upload!", file.Filename))
		//传到项目根目录，名字用本身
		// type1 := ctx.DefaultPostForm("type","alert")
		// username := ctx.PostForm("username")
		// password := ctx.PostForm("password")
		// hobby := ctx.PostFormArray("hobby")
		// ctx.String(http.StatusOK,fmt.Sprintf("type is %s, username is %s, password is %s. hobby is %v",
		// type1,username,password,hobby))
	})
	r.Run(":8000")
}
