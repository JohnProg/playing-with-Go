package models

import (
	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
	// users *mgo.Collection
	books *mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB
	// users = db.C("users")
	books = db.C("books")
}
