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
	organizations *mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB
	designers = db.C("designers")
	organizations = db.C("organizations")

	indexOrganizationName := mgo.Index{
		Key:    []string{"name", "ruc", "organizationId"},
		Unique: true,
	}
	organizations.EnsureIndex(indexOrganizationName)
}
