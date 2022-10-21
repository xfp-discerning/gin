package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if cookie, err := ctx.Cookie("root"); err == nil {
			if cookie == "1234" {
				ctx.Next()
				return
			}
		}
		//return error
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "err.Error()"})
		//if auth is failed, use ctx.abort to stop the following actions
		ctx.Abort()
	}
}

func main() {
	r := gin.Default()
	r.GET("login", func(ctx *gin.Context) {
		ctx.SetCookie("root", "1234", 60, "/", "localhost", false, true)
		//返回信息
		ctx.String(200, "login success")
	})

	r.GET("home", Auth(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "111"})
	})
	r.Run(":8000")
}
