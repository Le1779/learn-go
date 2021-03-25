package users

import (
	middleware "learn-go/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {

	router.POST("/login", Login)

	router.Use(middleware.JWT())
	{
		router.GET("/users", GetAllUsers)
		router.GET("/users/:id", GetUser)
		router.POST("/users", PostUser)
		router.PUT("/users", PutUser)
		router.DELETE("/users/:id", DeleteUser)
	}
}
