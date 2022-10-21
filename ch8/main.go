package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Mytime(c *gin.Context) {
	t := time.Now()
	c.Next()
	t1 := time.Since(t)
	fmt.Println("共用时：", t1)
}

func main() {
	r := gin.Default()
	r.Use(Mytime)
	shopgroup := r.Group("/shopping")
	{
		shopgroup.GET("/a", shopahander)
		shopgroup.GET("/b", shopbhander)
	}
	r.Run(":8000")
}

func shopahander(c *gin.Context) {
	c.JSON(200, gin.H{"xfp": "handsome"})
	time.Sleep(5 * time.Second)
}

func shopbhander(c *gin.Context) {
	time.Sleep(5 * time.Second)
}
