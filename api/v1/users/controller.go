package users

import (
	database "learn-go/database"
	jwt "learn-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token, err := jwt.GenerateToken(12)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": token,
	})
}

func GetUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}

	var user database.User = database.User{ID: id}
	var isExist bool
	isExist, err = user.ExistUser()
	if !isExist {
		c.JSON(200, gin.H{
			"message": "ID is not exist.",
		})
		return
	}

	var users []database.User
	users, err = user.QueryUser(0, 0)
	c.JSON(200, gin.H{
		"message": users,
	})
}

func GetAllUsers(c *gin.Context) {
	var user database.User = database.User{}

	users, err := user.QueryUser(0, 0)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": users,
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

func PutUser(c *gin.Context) {
	idStr := c.PostForm("ID")
	name := c.PostForm("Name")
	password := c.PostForm("Password")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "ID is not number.",
		})
		return
	}

	if name == "" || password == "" {
		c.JSON(200, gin.H{
			"message": "null params.",
		})
		return
	}

	var user database.User = database.User{ID: id, Name: name, Password: password}
	var isExist bool
	isExist, err = user.ExistUser()
	if !isExist {
		c.JSON(200, gin.H{
			"message": "ID is not exist.",
		})
		return
	}

	err = user.UpdateUser()
	message := "Updated user " + name
	if err != nil {
		message = err.Error()
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}

func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}

	var user database.User = database.User{ID: id}
	var isExist bool
	isExist, err = user.ExistUser()
	if !isExist {
		c.JSON(200, gin.H{
			"message": "ID is not exist.",
		})
		return
	}

	err = user.DeleteUser()
	message := "Deleted user."
	if err != nil {
		message = err.Error()
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}
