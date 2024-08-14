package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func postRegister(con *gin.Context) {

	var user User
	if err := con.ShouldBindJSON(&user); err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	con.JSON(http.StatusOK, gin.H{"data": user})
}

func main() {
	router := gin.Default()
	router.POST("/register", postRegister)
	router.Run(":8081")
}
