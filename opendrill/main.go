package main

import (
	"./app/models"
	"./router"
	"github.com/bradfitz/gomemcache/memcache"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	// "runtime"
)

const (
	MONGO_URLS    = "mongodb://127.0.0.1"
	DATABASE_NAME = "opendrill"
	MEMCACHE_URLS = "localhost:11211"
)

var (
	session *mgo.Session
	db      *mgo.Database
	mc      *memcache.Client
)

// // Connect to DB
func init() {
	var err error
	session, err = mgo.Dial(MONGO_URLS)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB '%s'", MONGO_URLS)
	}
	// session.SetMode(mgo.Monotonic, true)
	session.SetMode(mgo.Strong, true) // Most similar to Postgres
	db = session.DB(DATABASE_NAME)
	models.SetDB(db)
}

// // Connect to cache
func init() {
	mc = memcache.New(MEMCACHE_URLS)
	// handlers.SetCache(mc)
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	defer session.Close()

	router.Init()

	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
