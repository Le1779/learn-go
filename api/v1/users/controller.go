package users

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is Kevin Login.",
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Hello" + id,
	})
}
