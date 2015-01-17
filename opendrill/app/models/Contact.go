package models

import "gopkg.in/mgo.v2/bson"
import "time"

<<<<<<< HEAD
type Contact struct{
	Id 			bson.ObjectId 	`bson:"_id" json:"id"`
	Name 		string			`db:"name" json:"name"`
	Email 		string			`db:"email" json:"email"`
	Categories 	[]string 		`db:"categories" json:"categories"`
	ContactList []string 		`db:"contactlist" json:"contactlist"`
	Created     time.Time 		`db:"created" json:"created"`
	Modified    time.Time 		`db:"modified" json:"modified"`

=======
type Contact struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `db:"name" json:"name"`
	Email       string        `db:"email" json:"email"`
	Categories  []string      `db:"categories" json:"categories"`
	ContactList []string      `db:"contactlist" json:"contactlist"`
	CreatedAt   time.Time     `json:"createdAt"`
	ModifiedAt  time.Time     `json:"updatedAt"`
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada
}

func AllContact() (contact2 []Contact, err error) {
	err = contacts.Find(nil).All(&contact2)
	return
}

func GetContact(Id string) (err error, contact Contact) {
	bid := bson.ObjectIdHex(Id)
	err = contacts.
		FindId(bid).
		One(&contact)
	return
}

func CreateContact(contact Contact) (err error, contact2 Contact) {
	contact2 = contact
<<<<<<< HEAD
	contact2.Created = time.Now()
	contact2.Modified = time.Now()
=======
	contact2.CreatedAt = time.Now()
	contact2.ModifiedAt = time.Now()
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada
	contact2.Id = bson.NewObjectId()

	if err := contacts.Insert(contact2); err != nil {
		return err, contact
	}
	return nil, contact2
}

func RemoveContact(Id string) (err error, deleted bool) {
	deleted = false
	bid := bson.ObjectIdHex(Id)
	err = contacts.Remove(bson.M{"_id": bid})

	if err != nil {
		return err, deleted
	}
	deleted = true
	return nil, deleted
}

func UpdateContact(contact Contact, Id string) (err error, contact2 Contact) {
	contact2 = contact
	bid := bson.ObjectIdHex(Id)
	err = contacts.Update(bson.M{"_id": bid},
		bson.M{
<<<<<<< HEAD
			"name": contact2.Name,
			"email": contact2.Email,
			"categories": contact2.Categories,
			"contactlist": contact2.ContactList,
			"modified": time.Now(),
			"_id":    bid,
=======
			"name":        contact2.Name,
			"email":       contact2.Email,
			"categories":  contact2.Categories,
			"contactlist": contact2.ContactList,
			"createdAt":   contact2.CreatedAt,
			"modifiedAt":  time.Now(),
			"_id":         bid,
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada
		})
	if err != nil {
		return err, contact
	}
	contact2.Id = bid
	return nil, contact2
}
