package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Event struct {
	Id int64
	Name string
}

type User struct {
	Id int64
	UserName string
	UserEmail string
	UserPassword string
}

func main() {
	db, _ := gorm.Open("postgres", "user=postgres dbname=pqgotest password=1234 sslmode=disable")
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True")
	// db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	db.DB()
	db.DropTableIfExists(&Event{})
	db.CreateTable(&Event{})

	events := []Event{Event{Name: "Blink 182 Concert"}, Event{Name: "The Offspring Concert"}}
	events = append(events, Event{Name: "Metallica Concert"})
	events = append(events, Event{Name: "Linkin Park Concert"})
	events = append(events, Event{Name: "Avicii Concert"})
	events = append(events, Event{Name: "One Republic Concert"})

	for i:=0; i<len(events); i++ {
		db.Create(events[i])
	}

	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})

	user := User{UserName: "Jinzhu", UserEmail: "john.cfmr.2009@gmail.com", UserPassword: "123"}
	db.Create(&user)
}
