package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}

func main() {
	router := gin.Default()

	router.Run(":8080")
}
