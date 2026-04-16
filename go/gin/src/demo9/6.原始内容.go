package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func main() {
	r := gin.Default()
	r.POST("", func(c *gin.Context) {
		byteData, _ := io.ReadAll(c.Request.Body)
		fmt.Println(string(byteData))
		c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
		fmt.Println(c.Request.Header)
	})
	r.Run(":8080")
}
