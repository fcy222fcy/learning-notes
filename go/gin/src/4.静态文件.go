package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 前面的是别名
	r.Static("st", "static")
	r.Run(":8080")
}
