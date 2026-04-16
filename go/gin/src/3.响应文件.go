package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream") // 表示是文件流,浏览器可以直接下载
		c.Header("Content-Disposition", "attachment;filename=3.响应文件.go")
		c.File("3.响应文件1.go")
	})
	r.Run(":8080")
}
