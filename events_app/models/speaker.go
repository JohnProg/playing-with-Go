package data

import (
	"fmt"
)

type Speaker struct {
	Person     //type embedding for composition
	SpeaksOn   []string
	PastEvents []string
}

//overrides GetDetails
func (s Speaker) GetDetails() {
	//Call person GetDetails
	s.Person.GetDetails()
	fmt.Println("Speaker talks on following technologies:")
	for _, value := range s.SpeaksOn {
		fmt.Println(value)
	}
	fmt.Println("Presented on the following conferences:")
	for _, value := range s.PastEvents {
		fmt.Println(value)
	}
}
