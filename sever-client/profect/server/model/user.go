package model

//Define the struct of a user

type User struct {
	UserId string         `json:"user_id"`
	UserPwd string	   `json:"user_pwd"`
	UserName string    `json:"user_name"`
}