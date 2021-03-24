package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {

	router.POST("/login", Login)
	router.GET("/users/:id", GetUser)
	router.POST("/users", PostUser)
}
