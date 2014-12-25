package data

import (
	"fmt"
)

type Organizer struct {
	Person //type embedding for composition
	events []string
}

//overrides GetDetails
func (o Organizer) GetDetails() {
	//Call person GetDetails
	o.Person.GetDetails()
	fmt.Println("Organizer, conducting following Meetups:")
	for _, value := range o.meetups {
		fmt.Println(value)
	}
}