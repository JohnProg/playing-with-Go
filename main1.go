package main

import (
	"fmt"
)

type User struct {
	FirstName string
}

//var users map[string]User

func main() {
	//var fruits [3]string
	fruits := [3]string{}
	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[2] = "strawberry"
	//fruits[3] = "banana"
	for i:=0; i<len(fruits); i++ {
		fmt.Println(fruits[i])
	}
	fmt.Println(fruits[2:3])


	u1 := User{FirstName: "Chris"}
	u2 := User{FirstName: "Roger"}
	users := make(map[string]User)
	users["001"] = u1
	users["002"] = u2

	fmt.Println(users)
}
