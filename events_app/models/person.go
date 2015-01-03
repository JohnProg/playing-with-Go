package data

import (
	"fmt"
)

type People interface {
	SayHello()
	GetDetails()
}

type Person struct {
	Name        string
	Age         int
	City, Phone string
}

//A person method
func (p Person) SayHello() {
	fmt.Printf("Hi, I am %s, from %s\n", p.Name, p.City)
}

//A person method
func (p Person) GetDetails() {
	fmt.Printf("[Name: %s, Age: %d, City: %s, Phone: %s]\n", p.Name, p.Age, p.City, p.Phone)
}
