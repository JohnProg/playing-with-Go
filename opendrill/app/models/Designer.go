package models

import "gopkg.in/mgo.v2/bson"

type Designer struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	User      `bson:"user" json:"user"`
	Templates []Template `bson:"templates" json:"templates"`
}

func AllDesigners() (designers2 []Designer, err error) {
	err = designers.
		Find(nil).
		All(&designers2)
	return
}

func CreateDesigner(designer Designer) (err error, designer2 Designer) {
	designer2 = designer
	designer2.Id = bson.NewObjectId()
	designer2.Templates = []Template{}

	if err := designers.Insert(designer2); err != nil {
		return err, designer
	}
	return nil, designer2
}

func GetDesigner(designerID string) (err error, designer Designer) {
	bid := bson.ObjectIdHex(designerID)
	err = designers.
		FindId(bid).
		One(&designer)
	return
}

func RemoveDesigner(Id string) (err error, deleted bool) {
	deleted = false
	bid := bson.ObjectIdHex(Id)
	err = designers.
		Remove(bson.M{"_id": bid})
	if err != nil {
		return err, deleted
	}
	deleted = true
	return nil, deleted
}

func UpdateDesigner(designer Designer, Id string) (err error, designer2 Designer) {
	designer2 = designer
	bid := bson.ObjectIdHex(Id)
	err = designers.Update(bson.M{"_id": bid},
		bson.M{"user": designer2.User,
			"templates": designer2.Templates,
			"_id":       bid,
		})
	if err != nil {
		return err, designer
	}
	designer2.Id = bid
	return nil, designer2
}
