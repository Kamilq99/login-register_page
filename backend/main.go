package main

import (
	"login_register/register"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/register", register.Register)

	r.Run(":8080")
}
