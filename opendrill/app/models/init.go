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
	db           *mgo.Database
	designers    *mgo.Collection
	organizators *mgo.Collection
	templates    *mgo.Collection
	contacts     *mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB
	designers = db.C("designers")
	templates = db.C("templates")
	contacts = db.C("contacts")
	organizators = db.C("organizators")

	contacts.EnsureIndexKey("email")

	index := mgo.Index{
		Key:    []string{"user.username"},
		Unique: true,
	}
	designers.EnsureIndex(index)
	organizators.EnsureIndex(index)

	index2 := mgo.Index{
		Key:    []string{"name", "designerId"},
		Unique: true,
	}
	templates.EnsureIndex(index2)
}
