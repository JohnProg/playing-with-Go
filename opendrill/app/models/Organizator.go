package models

import "gopkg.in/mgo.v2/bson"

type Organizator struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	User
	Organizations []Organization `json:"organizations"`
}
