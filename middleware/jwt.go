package middleware

import (
	utils "learn-go/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		token := authHeader[len(BEARER_SCHEMA):]

		if token == "" {
			c.JSON(200, gin.H{
				"message": "check token fail1",
			})

			c.Abort()
			return
		}

		_, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "check token fail2",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
