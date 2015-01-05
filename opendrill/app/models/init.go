package models

import (
	"gopkg.in/mgo.v2"
)

// error response contains everything we need to use http.Error
type HandlerError struct {
	Error   error
	Message string
	Code    int
}

var (
	db *mgo.Database
	books *mgo.Collection
	contacts *mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB
	// users = db.C("users")
	books = db.C("books")
	contacts = db.C("contacts")
}
