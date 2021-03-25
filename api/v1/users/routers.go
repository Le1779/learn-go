package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {

	router.POST("/login", Login)
	router.GET("/users/:id", GetUser)
	router.GET("/users", GetAllUsers)
	router.POST("/users", PostUser)
	router.PUT("/users", PutUser)
	router.DELETE("/users/:id", DeleteUser)
}
