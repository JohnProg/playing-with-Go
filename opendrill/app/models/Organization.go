package models

import "gopkg.in/mgo.v2/bson"

type Organization struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Name        string
	Description string
	Logo        string
	RUC         string
	MandrillKey string
}
