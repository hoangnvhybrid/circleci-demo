package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	)

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello,World!")
	})

	r.GET("/welcome", func(c *gin.Context) {
		c.String(200, "Page welcome!!!")
	})

	r.Run(":8086")

}
