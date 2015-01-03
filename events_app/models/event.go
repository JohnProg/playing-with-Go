package data

import (
	"time"
)

type Event struct {
	Name     string
	Location string
	City     string
	Date     time.Time
	People2  []People // speaker, organizator, attendee
}

func (e Event) MeetupPeople() {
	for _, v := range e.People2 {
		v.SayHello()
		v.GetDetails()
	}
}
