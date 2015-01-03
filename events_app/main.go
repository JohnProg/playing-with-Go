package main

import (
	models "./models"
	"time"
)

func main() {
	//Load data
	shiju := models.Speaker{
		models.Person{"Shiju", 35, "Kochi", "+91-94003372xx"},
		[]string{"Go", "Docker", "Azure", "AWS"},
		[]string{"FOSS", "JSFOO", "MS TechDays"},
	}
	satish := models.Organizer{
		models.Person{"Satish", 35, "Pune", "+91-94003372xx"},
		[]string{"Gophercon", "RubyConf"},
	}
	alex := models.Attendee{
		models.Person{"Alex", 22, "Bangalore", "+91-94003672xx"},
		[]string{"Go", "Ruby"},
	}

	event := models.Event{
		"Evento de html5",
		"av. los alamos S.J.M",
		"Lima",
		time.Date(2015, time.January, 19, 9, 0, 0, 0, time.UTC),
		[]models.People{shiju, satish, alex},
	}

	//get details of meetup people
	event.MeetupPeople()
}
