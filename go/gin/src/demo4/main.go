package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		//魏正想到此一游.
		c.String(200, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}
