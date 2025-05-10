package login

import (
	"login_register/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User_Login
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(400, gin.H{
			"message": "Please fill all the details",
		})
		return
	}

	registeredUser, exists := models.User_Registers[user.Username]
	if !exists || registeredUser.Password != user.Password {
		c.JSON(401, gin.H{
			"message": "Login failed, Username or Password is incorrect",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
	})
}
