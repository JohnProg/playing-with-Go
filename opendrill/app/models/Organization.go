package models

import (
	"gopkg.in/mgo.v2/bson"
	//e "../error"
	"time"
)

type Organization struct {
	Id            bson.ObjectId `bson:"_id" json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Logo          string        `json:"logo"`
	RUC           string        `json:"ruc"`
	MandrillKey   string        `json:"mandrillKey"`

	Users 		  []User 		`bson:"users" json:"users"`
	ListContacts  []ListContact `bson:"listContacts" json:"listContacts"`
	Templates	  []Template 	`bson:"templates" json:"templates"`	

	CreatedAt     time.Time     `json:"createdAt"`
	ModifiedAt    time.Time     `json:"updatedAt"`
}

//Commons

func OrganizationFindId(id string) (err error, result Organization) {
	err = organizations.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func AllOrganizations() (organizations2 []Organization, err error) {
	err = organizations.Find(nil).All(&organizations2)
	return
}

//Methods

func AddOrganizationFromOrganizator(user User, organization Organization) (err error, organization2 Organization) {
	var new_organization Organization
	_user := []User{user}
	_template := []Template{}
	_list_conatct := []ListContact{}

	organization2 = new_organization
	organization2.Users = _user
	organization2.ListContacts = _list_conatct
	organization2.Templates = _template
	organization2.Name = organization.Name
	organization2.Description = organization.Description
	organization2.RUC = organization.RUC
	organization2.CreatedAt = time.Now()
	organization2.ModifiedAt = time.Now()
	organization2.Id = bson.NewObjectId()
	if err := organizations.Insert(organization2); err != nil {
		return err, new_organization
	}
	return nil, organization2
}


func GetOrganizationsFromOrganizator(organizatorID string) (err error, organizations2 Organization) {

	err = organizations.Find(nil).
		  Select(bson.M{"users": bson.M{"$elemMatch": bson.M{"_id":  bson.ObjectIdHex(organizatorID)}}}).
		  One(&organizations2)

	return
}

func AddUserToOrganization(user User, organizationID string) (err error, organization Organization) {
	/*
		db.organizations.update( {_id: ObjectId("54c3e639b71b7f1fed000002")}, {$push:{listContacts: "prueba"}} )
	*/
	upsert_user := bson.M{ "users": user}
	upsert_push := bson.M{ "$push": upsert_user}

	err = organizations.Update(bson.M{"_id": bson.ObjectIdHex(organizationID)},upsert_push)
	if err != nil {
		return err, organization
	}
	err, organization = OrganizationFindId(organizationID)
	if err != nil {
		return err, organization
	}

	return nil, organization
}

func DeleteUserToOrganization(userID string, organizationID string) (err error, organization Organization) {
	upsert_condition 	:= bson.M{ "_id": bson.ObjectIdHex(userID)}
	upsert_user 		:= bson.M{ "users": upsert_condition}
	upsert_pull	  		:= bson.M{ "$pull": upsert_user}

	err = organizations.Update(bson.M{"_id": bson.ObjectIdHex(organizationID)},upsert_pull)
	if err != nil {
		return err, organization
	}

	err, organization = OrganizationFindId(organizationID)
	if err != nil {
		return err, organization
	}

	return nil, organization
}
