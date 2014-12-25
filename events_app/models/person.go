package data

import (
	"fmt"
)

type People interface {
	SayHello()
	GetDetails()
}

type Person struct {
	name string
	age int
	city,phone string
}

//A person method
func (p Person ) SayHello() {
	fmt.Printf("Hi, I am %s, from %s\n", p.name, p.city)
}

//A person method
func (p Person) GetDetails() {
	fmt.Printf("[Name: %s, Age: %d, City: %s, Phone: %s]\n", p.name,p.age, p.city, p.phone)
}