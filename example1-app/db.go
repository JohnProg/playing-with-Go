package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Event struct {
	Id int64
	Name string
}

func main() {
	db, _ := gorm.Open("postgres", "user=chris dbname=higo sslmode=disable")
	db.DB()
	//db.SingularTable(true)
	//db.DropTable(&Event{})
	//db.CreateTable(&Event{})
	//log.Println(&event)
	//db.NewRecord(event)

	events := []Event{Event{Name: "Blink 182 Concert"}, Event{Name: "The Offspring Concert"}}
	//events = append(events, Event{Name: "Metallica Concert"})
	//events = append(events, Event{Name: "Linkin Park Concert"})
	//events = append(events, Event{Name: "Avicii Concert"})
	//events = append(events, Event{Name: "One Republic Concert"})

	for i:=0; i<len(events); i++ {
		db.Create(events[i])
	}
}
