package users

import (
	database "learn-go/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is Kevin Login.",
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}

	var user database.User = database.User{}
	user.QueryUserById(i)
	c.JSON(200, gin.H{
		"message": user,
	})
}

func PostUser(c *gin.Context) {
	name := c.PostForm("Name")
	password := c.PostForm("Password")

	if name == "" || password == "" {
		c.JSON(200, gin.H{
			"message": "null params.",
		})
		return
	}

	var user database.User = database.User{Name: name, Password: password}
	err := user.CreateUser()

	message := "Created user " + name
	if err != nil {
		message = err.Error()
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}
