package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.Run(":8086")

}
