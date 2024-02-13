package entity

type User struct {
	Name   string
	Age    int
	Friend []int
}

func NewUser() *User {
	var user User
	return &user
}
