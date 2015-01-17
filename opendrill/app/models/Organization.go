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
	OrganizatorID bson.ObjectId `json:"organizatorId" bson:"organizatorId"`
	MandrillKey   string        `json:"mandrillKey"`
	CreatedAt     time.Time     `json:"createdAt"`
	ModifiedAt    time.Time     `json:"updatedAt"`
}

func GetOrganizationsFromOrganizator(organizatorID string) (organizations2 []Organization) {
	organizations.
		Find(bson.M{"organizatorId": bson.ObjectIdHex(organizatorID)}).
		All(&organizations2)
	return
}

func GetOrganizationFromOrganizator(organizatorID string, organizationID string) (err error, organization2 Organization) {
	err = organizations.
		Find(bson.M{"organizatorId": bson.ObjectIdHex(organizatorID), "_id": bson.ObjectIdHex(organizationID)}).
		One(&organization2)
	return
}

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
