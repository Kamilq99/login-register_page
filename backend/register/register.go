package register

import (
	"login_register/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var user models.User_Register
	c.BindJSON(&user)

	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(400, gin.H{
			"message": "Please fill all the details",
		})
		return
	}

	for _, v := range models.User_Registers {
		if v.Username == user.Username {
			c.JSON(400, gin.H{
				"message": "User already exists",
			})
			return
		}
	}

	for _, v := range models.User_Registers {
		if v.Email == user.Email {
			c.JSON(400, gin.H{
				"message": "Email already exists",
			})
			return
		}
	}

	models.User_Registers[user.Username] = user

	c.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}
