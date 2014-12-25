package main

import (
	"fmt"
	"time"
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
type Speaker struct {
	Person //type embedding for composition
	speaksOn []string
	pastEvents []string
}
//overrides GetDetails
func (s Speaker) GetDetails() {
	//Call person GetDetails
	s.Person.GetDetails()
	fmt.Println("Speaker talks on following technologies:")
	for _, value := range s.speaksOn {
		fmt.Println(value)
	}
	fmt.Println("Presented on the following conferences:")
	for _, value := range s.pastEvents {
		fmt.Println(value)
	}
}
type Organizer struct {
	Person //type embedding for composition
	events []string
}
//overrides GetDetails
func (o Organizer) GetDetails() {
	//Call person GetDetails
	o.Person.GetDetails()
	fmt.Println("Organizer, conducting following Meetups:")
	for _, value := range o.events {
		fmt.Println(value)
	}
}
type Attendee struct {
Person //type embedding for composition
interests []string
}
type Event struct
{
	name string
	location string
	city string
	date time.Time
	people []People // speaker, organizator, attendee
}
func (e Event) MeetupPeople(){
	for _, v := range e.people {
		v.SayHello()
		v.GetDetails()
	}
}