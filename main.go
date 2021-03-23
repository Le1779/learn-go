package main

import (
	router "learn-go/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := router.InitRouter()

	server.LoadHTMLGlob("frontend/*")

	server.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello Gin",
		})
	})

	server.Run(":8888")
}
