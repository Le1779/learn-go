package main

import (
	database "learn-go/database"
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
	database.Setup()

	server.Run(":8888")
}
