package api

import (
	v1 "learn-go/api/v1"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {
	api := router.Group("/api")
	{
		v1.RegisterRouter(api)
	}
}
