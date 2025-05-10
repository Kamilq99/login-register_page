package models

type User_Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User_Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var User_Registers = make(map[string]User_Register)
