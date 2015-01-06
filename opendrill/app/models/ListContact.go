package models

import  "gopkg.in/mgo.v2/bson"
import "time"


type ListContact struct{
	Id 			bson.ObjectId  `bson:" json:"id"`
	Name 		string			
	Status 		bool
	Detail 	    []DetailContact `bson:"detail"`
	Created     time.Time 		`db:"created" json:"created"`
	Modified    time.Time 		`db:"modified" json:"modified"`
}

type DetailContact struct {
	Contact      Contact `bson:"contact"`
}

func AllListContact() (list_contact2 []ListContact, err error){
	err = list_contacts.Find(nil).All(&list_contact2)
	return
}

func CreateListContact(list_contact ListContact) (err error, list_contact2 ListContact) {
	list_contact2 = list_contact
	list_contact2.Created = time.Now()
	list_contact2.Modified = time.Now()
	list_contact2.Id = bson.NewObjectId()

	if err := list_contacts.Insert(list_contact2); err != nil {
		return err, list_contact
	}
	return nil, list_contact2

}