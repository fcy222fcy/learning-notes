package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/*")
	r.GET("", func(c *gin.Context) {
		// obj是文件接收的参数
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
