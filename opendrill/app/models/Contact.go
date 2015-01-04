package models

import "gopkg.in/mgo.v2/bson"

type Contact struct{
	Id bson.ObjectId `bson:"_id" json:"id"`
	Name string
	Email string
}

func AllContact() (contact2 []Book, err error) {
	err = contacts.Find(nil).All(&contact2)
	return
}
