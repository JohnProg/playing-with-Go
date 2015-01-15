package models

import "gopkg.in/mgo.v2/bson"
import "time"

type Category struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	CreatedAt  time.Time     `json:"createdAt"`
	ModifiedAt time.Time     `json:"updatedAt"`
}

func AllCategory() (category2 []Category, err error) {
	err = categories.Find(nil).All(&category2)
	return
}

func GetCategory(Id string) (err error, category Category) {
	bid := bson.ObjectIdHex(Id)
	err = categories.
		FindId(bid).
		One(&category)
	return
}

func CreateCategory(category Category) (err error, category2 Category) {
	category2 = category
	category2.CreatedAt = time.Now()
	category2.ModifiedAt = time.Now()
	category2.Id = bson.NewObjectId()

	if err := categories.Insert(category2); err != nil {
		return err, category
	}
	return nil, category2
}

func RemoveCategory(Id string) (err error, deleted bool) {
	deleted = false
	bid := bson.ObjectIdHex(Id)
	err = categories.Remove(bson.M{"_id": bid})

	if err != nil {
		return err, deleted
	}
	deleted = true
	return nil, deleted
}

func UpdateCategory(category Category, Id string) (err error, category2 Category) {
	category2 = category
	bid := bson.ObjectIdHex(Id)
	err = categories.Update(bson.M{"_id": bid},
		bson.M{
			"name":      category2.Name,
			"updatedAt": time.Now(),
			"createdAt": category2.CreatedAt,
			"_id":       bid,
		})
	if err != nil {
		return err, category
	}
	category2.Id = bid
	return nil, category2
}
