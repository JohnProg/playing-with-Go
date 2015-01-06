package main

import (
	"gopkg.in/mgo.v2"
	"flag"
	"log"
	"fmt"
	"net/http"
	"./app/models"
	"./router"
)

const (
	MONGO_URLS    = "mongodb://127.0.0.1"
	DATABASE_NAME = "opendrill"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

func init() {
	var err error
	session, err = mgo.Dial(MONGO_URLS)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB '%s'", MONGO_URLS)
	}
	session.SetMode(mgo.Strong, true)
	db = session.DB(DATABASE_NAME)
	models.SetDB(db)
}

func CloseSession() {
	session.Close()
}

func main() {
	port := flag.Int("port", 6969, "port to serve on")
	flag.Parse()

	defer session.Close()

	router.Init()

	log.Printf("Running on port %d\n", *port)
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	err := http.ListenAndServe(addr, nil)
	log.Println(err.Error())
}
