package message

//Define the struct of a user

type User struct {
	Id int         `json:"id"`
	Pwd string	   `json:"pwd"`
	Name string    `json:"name"`
}