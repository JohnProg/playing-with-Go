package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db, _ = gorm.Open("postgres", "user=postgres dbname=higo password=1234 sslmode=disable")

type Event struct {
	Id int64
	Name string
}

func availableEvents() (events []Event) {
	//events = append(events, Event{Name: "Godsmack Concert"})
	//events = append(events, Event{Name: "Linkin Park Concert"})
	events = db.Find(&Event)
	return events
}

func eventCreator(request *http.Request) (event Event, error error) {
	return Event{Name: "abc"}, nil
}

func eventsHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		log.Println("GET")
		events, _ := json.Marshal(availableEvents())
		response.Write(events)
	} else if request.Method == "POST" {
		log.Println("POST")
		event, _ := eventCreator(request)
		eventJson, _ := json.Marshal(event)
		response.Write(eventJson)
	}
}

func main() {

	http.HandleFunc("/events", eventsHandler)
	http.ListenAndServe(":8080", nil)
}


