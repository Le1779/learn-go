package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

    server.LoadHTMLGlob("frontend/*")
	server.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello Gin",
		})
	})

	server.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is Kevin.",
		})
	})

	server.Run(":8888")
}