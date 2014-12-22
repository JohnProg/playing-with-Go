package main

import (
	"fmt"
	"reflect"
)

type User struct {
	FirstName string
	LastName string
}

func main() {
	user := User{}
	userr := new(User)
	userPointer := &user
	userPointer.FirstName = "Roger"
	fmt.Println(reflect.TypeOf(user))
	fmt.Println(reflect.TypeOf(userr))
	fmt.Println(reflect.TypeOf(userPointer))
	fmt.Println(user.FirstName)
	fmt.Println(userPointer.FirstName)
	fmt.Println(userr.FirstName)
}
