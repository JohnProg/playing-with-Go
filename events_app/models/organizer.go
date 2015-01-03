package data

import (
	"fmt"
)

type Organizer struct {
	Person //type embedding for composition
	Events []string
}

//overrides GetDetails
func (o Organizer) GetDetails() {
	//Call person GetDetails
	o.Person.GetDetails()
	fmt.Println("Organizer, conducting following Meetups:")
	for _, value := range o.Events {
		fmt.Println(value)
	}
}
