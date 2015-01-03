package main

import (
	"./router"
	"fmt"
	// "github.com/bradfitz/gomemcache/memcache"
	// "gopkg.in/mgo.v2"
	"flag"
	"log"
	"net/http"
	// "runtime"
)

// const (
// 	MONGO_URLS    = "mongodb://127.0.0.1"
// 	DATABASE_NAME = "opendrill"
// 	MEMCACHE_URLS = "localhost:11211"
// )

// var (
// 	session *mgo.Session
// 	db      *mgo.Database
// 	mc      *memcache.Client
// )

// // Connect to DB
// func init() {
// 	var err error
// 	session, err = mgo.Dial(MONGO_URLS)
// 	if err != nil {
// 		log.Fatalf("Error connecting to MongoDB '%s'", MONGO_URLS)
// 	}
// 	// session.SetMode(mgo.Monotonic, true)
// 	session.SetMode(mgo.Strong, true) // Most similar to Postgres
// 	db = session.DB(DATABASE_NAME)
// 	// TODO: It doesn't seem like you should have to do this...
// 	// Tell other packages which Mongo database to use
// 	helpers.SetDB(db)
// 	types.SetDB(db)
// }

// // Connect to cache
// func init() {
// 	mc = memcache.New(MEMCACHE_URLS)
// 	handlers.SetCache(mc)
// }

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// defer session.Close()

	// command line flags
	port := flag.Int("port", 8001, "port to serve on")
	flag.Parse()

	router.Init()

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
