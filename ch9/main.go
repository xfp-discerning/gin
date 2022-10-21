package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("cookie", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("key")
		if err != nil {
			cookie = "NotSet"
			//name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool
			ctx.SetCookie("key", "1234", 60, "/", "localhost", false, true)
		}
		fmt.Println(err)
		fmt.Println("the value of cookie is ", cookie)
	})
	r.Run(":8000")
}
