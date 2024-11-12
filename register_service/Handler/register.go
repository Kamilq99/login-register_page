package handler

import (
	"net/http"
	"register_service/db_connect"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RegisterHanlder(w http.ResponseWriter, r *http.Request) {

	username := "root"
	password := "<password>"
	dbadress := "127.0.0.1:3306"
	dbname := "login_register"

	db_connect.DBconnect(username, password, dbadress, dbname)

	var user User
}
