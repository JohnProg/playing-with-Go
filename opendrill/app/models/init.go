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
	db            *mgo.Database
	designers     *mgo.Collection
	organizators  *mgo.Collection
	organizations *mgo.Collection
	templates     *mgo.Collection
	contacts      *mgo.Collection
	list_contacts *mgo.Collection
	categories    *mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB
	designers = db.C("designers")
	templates = db.C("templates")
	organizators = db.C("organizators")
	organizations = db.C("organizations")
	contacts = db.C("contacts")
	list_contacts = db.C("list_contacts")
	categories = db.C("categories")

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

	index3 := mgo.Index{
		Key:    []string{"name", "organizatorId"},
		Unique: true,
	}
	organizations.EnsureIndex(index3)
}
