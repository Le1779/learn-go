package routers

import (
	api "learn-go/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	r := router.Group("")
	{
		api.RegisterRouter(r)
	}

	return router
}
