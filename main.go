package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// This handler will match /user/john but will not match /user/ or /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// Multipart/Urlencoded Form POST request
	r.POST("/form", func(c *gin.Context) {
		message := c.PostForm("message")
		user := c.DefaultPostForm("user", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"user":    user,
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}