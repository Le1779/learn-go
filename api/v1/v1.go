package v1

import (
	users "learn-go/api/v1/users"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		users.RegisterRouter(v1)
	}
}
