package main

import (
	"./models"
	"./routes"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

const (
	MONGO_URLS    = "mongodb://127.0.0.1"
	DATABASE_NAME = "BibliotecaVirtual"
)

func main() {
	var err error
	session, err = mgo.Dial(MONGO_URLS) //os.Getenv("MONGO_URLS")
	if err != nil {
		log.Fatalf("Error connecting to MongoDB '%s'", MONGO_URLS)
	}
	session.SetMode(mgo.Strong, true) // Strong or Monotonic
	db = session.DB("BibliotecaVirtual")
	models.SetDB(db)

	defer session.Close()
	router.Init()
	log.Println("Listening on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err.Error())
	}
}
