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
