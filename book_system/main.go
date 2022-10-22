package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	//加载页面
	r.LoadHTMLGlob("template/*")
	//查看所有图书
	r.GET("/book/list", bookListHandler)
	r.Run(":8000")
}

func bookListHandler(c *gin.Context) {
	bookList, err := QueryAllbook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}
