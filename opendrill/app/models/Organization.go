package models

import (
	"gopkg.in/mgo.v2/bson"
	"log"
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


func AddOrganizationFromOrganizator(user User, organization Organization) (err error, organization2 Organization) {
	var new_organization Organization
	//user.Id = bson.NewObjectId()
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

func AllOrganizations() (organizations2 []Organization, err error) {
	err = organizations.Find(nil).All(&organizations2)
	return
}

func GetOrganization(organizationID string) (err error, organization2 Organization) {

	bid := bson.ObjectIdHex(organizationID)
	err = organizations.
		FindId(bid).
		One(&organization2)
	return
}


func GetOrganizationsFromOrganizator(organizatorID string) (err error, organizations2 Organization) {

	err = organizations.Find(nil).
		  Select(bson.M{"users": bson.M{"$elemMatch": bson.M{"_id":  bson.ObjectIdHex(organizatorID)}}}).
		  One(&organizations2)

	return
}

func AddUserToOrganization(user User, organizationID string) (err error, organization2 Organization) {
	log.Println("id:", organizationID)
	if err, organization2 := GetOrganization(organizationID); err != nil {
		return err, organization2
	}
	log.Println(organization2)
	/*var organizator Organizator
	if err, organizator := GetOrganizator(organizatorID); err != nil {
		log.Println(organizator)
		return err, Organization{}
	}
	user = User
	user.Id = bson.NewObjectId()
	if err := organizations.Insert(organization2); err != nil {
		return err, organization
	}

	organizators.Update(bson.M{"_id": organizator.Id},
		bson.M{
			"organizations": GetOrganizationsFromOrganizator(organizatorID),
		})*/
	return nil, organization2
}

/*
func AddOrganizationToOrganizator(organization Organization, organizatorID string) (err error, organization2 Organization) {
	var organizator Organizator
	if err, organizator := GetOrganizator(organizatorID); err != nil {
		log.Println(organizator)
		return err, Organization{}
	}
	organization2 = organization
	organization2.CreatedAt = time.Now()
	organization2.ModifiedAt = time.Now()
	organization2.Id = bson.NewObjectId()
	organization2.OrganizatorID = organizator.Id
	if err := organizations.Insert(organization2); err != nil {
		return err, organization
	}

	organizators.Update(bson.M{"_id": organizator.Id},
		bson.M{
			"organizations": GetOrganizationsFromOrganizator(organizatorID),
		})
	return nil, organization2
}

func RemoveOrganizationFromOrganizator(organizatorID string, organizationID string) (err error, deleted bool) {
	var organizator Organizator
	if err, organizator := GetOrganizator(organizatorID); err != nil {
		log.Println(organizator)
		return err, false
	}

	err = organizations.
		Remove(bson.M{"_id": bson.ObjectIdHex(organizationID)})
	if err != nil {
		return err, false
	}
	organizators.Update(bson.M{"_id": organizator.Id},
		bson.M{
			"organizations": GetOrganizationsFromOrganizator(organizatorID),
		})
	return nil, true
}

func UpdateOrganizationFromOrganizator(organization Organization, organizatorID string, organizationID string) (err error, organization2 Organization) {
	var organizator Organizator
	if err, organizator := GetOrganizator(organizatorID); err != nil {
		log.Println(organizator)
		return err, Organization{}
	}

	organization2 = organization
	bid := bson.ObjectIdHex(organizationID)
	err = organizations.Update(bson.M{"_id": bid},
		bson.M{"name": organization2.Name,
			"description":   organization2.Description,
			"logo":          organization2.Logo,
			"ruc":           organization2.RUC,
			"mandrillKey":   organization2.MandrillKey,
			"_id":           bid,
			"organizatorId": bson.ObjectIdHex(organizatorID),
			"createdAt":     organization2.CreatedAt,
			"modifiedAt":    time.Now(),
		})
	if err != nil {
		return err, organization
	}

	organizators.Update(bson.M{"_id": organizator.Id},
		bson.M{
			"organizations": GetOrganizationsFromOrganizator(organizatorID),
		})
	return nil, organization2
}
*/