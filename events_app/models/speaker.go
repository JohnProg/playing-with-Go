package data

import (
	"fmt"
)

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