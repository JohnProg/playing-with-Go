package models

import "gopkg.in/mgo.v2/bson"
import "log"

type Organizator struct {
	Id            bson.ObjectId `bson:"_id" json:"id"`
	User          `bson:"user" json:"user"`
	Organizations []Organization `bson:"organizations" json:"organizations"`
}

func GetOrganizator(organizatorID string) (err error, user_organizator User) {
	var organization Organization
	var users []User

	err = organizations.Find(nil).
		  Select(bson.M{"users": bson.M{"$elemMatch": bson.M{"_id":  bson.ObjectIdHex(organizatorID)}}}).
		  One(&organization)

    if err != nil{
    	return err, user_organizator
    }

    users = organization.Users
    for key := range users {
    	if users[key].Role == ROLE_ORGANIZATOR{
    		user_organizator = users[key]
    	}
    }

	return nil, user_organizator
}

func AddOrganization(organizatorID string, organization Organization)(err error, organization2 Organization){
	var user_organization User

	err, user_organization = GetOrganizator(organizatorID)
	if err != nil{
		return err, organization2
	}
	log.Println(user_organization)
	err, organization2 = AddOrganizationFromOrganizator(user_organization, organization)
	if err != nil{
		return err, organization2
	}
	return nil, organization2
} 

/*
func AllOrganizators(organizator Organizator) (organizators2 []Organizator, err error) {
	err = organizators.
		Find(nil).
		All(&organizators2)
	return
}

func CreateOrganizator(organizator Organizator) (err error, organizator2 Organizator) {
	organizator2 = organizator
	organizator2.Id = bson.NewObjectId()
	organizator2.Organizations = []Organization{}
	organizator2.User.Role = 0

	if err := organizators.Insert(organizator2); err != nil {
		return err, organizator
	}
	return nil, organizator2
}

func GetOrganizator(organizatorID string) (err error, organizator Organizator) {
	err = organizators.
		FindId(bson.ObjectIdHex(organizatorID)).
		One(&organizator)
	return
}

func RemoveOrganizator(organizatorID string) (err error, deleted bool) {
	deleted = false
	err = designers.
		Remove(bson.M{"_id": bson.ObjectIdHex(organizatorID)})
	if err != nil {
		return err, deleted
	}
	deleted = true
	return nil, deleted
}

func UpdateOrganizator(organizator Organizator, organizatorID string) (err error, organizator2 Organizator) {
	organizator2 = organizator
	bid := bson.ObjectIdHex(organizatorID)
	err = designers.Update(bson.M{"_id": bid},
		bson.M{"user": organizator2.User,
			"organizations": organizator2.Organizations,
			"_id":           bid,
		})
	if err != nil {
		return err, organizator
	}
	organizator2.Id = bid
	return nil, organizator2
}
*/