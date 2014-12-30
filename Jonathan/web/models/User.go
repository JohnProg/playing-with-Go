package models

type User struct{
	Id int
	Name string
	password string
}

func NewUser(name, password string) *User{
	return &User{name, password}
}