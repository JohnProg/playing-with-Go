package data

import (
	"time"
)

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
