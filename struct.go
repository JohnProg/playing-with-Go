package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (this Person) GetAllAsString() string {
	return this.FirstName + " " + this.LastName + " " + strconv.Itoa(this.Age)
}

func main() {
	var guy Person
	guy.FirstName = "John"
	guy.LastName = "Machahuay"
	guy.Age = 20

	//Print values
	fmt.Println(guy.FirstName, guy.LastName, guy.Age)

	// using the method
	fmt.Println(guy.GetAllAsString())
}
