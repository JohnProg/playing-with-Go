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
	db    *mgo.Database
	books *mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB
	addIndexToBookNames(db)
	books = db.C("books")
}

func addIndexToBookNames(mgoDB *mgo.Database) {
	db = mgoDB
	index := mgo.Index{
		Key:    []string{"title"},
		Unique: true,
	}
	indexErr := db.C("books").EnsureIndex(index)
	if indexErr != nil {
		panic(indexErr)
	}
}
