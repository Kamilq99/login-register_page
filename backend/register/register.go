package register

import (
	"login_register/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User_Register
	c.BindJSON(&user)
	c.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}
