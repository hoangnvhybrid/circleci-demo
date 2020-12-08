package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	// heroku se dung cong mac dinh, nen phai dung lenh duoi de lay, khi chay local thi nho set lai port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8086"
	}
	r.Run(":" + port)

}
