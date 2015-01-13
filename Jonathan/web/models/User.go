package models

type User struct{
	Id string
	Email string
	Password string
}

func NewUser(id, email, password string) *User{
	return &User{id, email, password}
}