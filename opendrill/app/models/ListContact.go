package models

import "gopkg.in/mgo.v2/bson"
import "time"

type ListContact struct {
	Id         bson.ObjectId 	`bson:" json:"id"`
	Name       string			`json:"name"`
	Status     bool				`json:"status"`
	Categories []string  		`bson:"categories" json:"categories"`
	CreatedAt  time.Time 		`json:"createdAt"`
	ModifiedAt time.Time 		`json:"updatedAt"`
}

func AllListContact() (list_contact2 []ListContact, err error) {
	err = list_contacts.Find(nil).All(&list_contact2)
	return
}

func CreateListContact(list_contact ListContact) (err error, list_contact2 ListContact) {
	list_contact2 = list_contact
	list_contact2.CreatedAt = time.Now()
	list_contact2.ModifiedAt = time.Now()
	list_contact2.Status = true
	list_contact2.Id = bson.NewObjectId()

	if err := list_contacts.Insert(list_contact2); err != nil {
		return err, list_contact
	}
	return nil, list_contact2

}

func GetListContact(Id string) (err error, list_contact ListContact) {
	bid := bson.ObjectIdHex(Id)
	err = list_contacts.
		FindId(bid).
		One(&list_contact)
	return
}

func RemoveListContact(Id string) (err error, deleted bool) {
	deleted = false
	bid := bson.ObjectIdHex(Id)
	err = list_contacts.
		Remove(bson.M{"_id": bid})
	if err != nil {
		return err, deleted
	}
	deleted = true
	return nil, deleted
}

func UpdateListContact(list_contact ListContact, Id string) (err error, list_contact2 ListContact) {
	list_contact2 = list_contact
	bid := bson.ObjectIdHex(Id)
	err = list_contacts.Update(bson.M{"_id": bid},
		bson.M{
			"name":       list_contact2.Name,
			"status":     list_contact2.Status,
			"categories": list_contact2.Categories,
			"createdAt":  list_contact2.CreatedAt,
			"modifiedAt": time.Now(),
			"_id":        bid,
		})
	if err != nil {
		return err, list_contact
	}
	list_contact2.Id = bid
	return nil, list_contact2
}
