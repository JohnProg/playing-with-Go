package client

import (
	models "../../models"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func GetOrganizationsFromOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	if !bson.IsObjectIdHex(organizatorID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	err, organizations := models.GetOrganizationsFromOrganizator(organizatorID)
	if err != nil {
		return []models.Organization{}, nil
	}
	return organizations, nil
}

func GetOrganization(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizationID := mux.Vars(r)["organizationID"]
	if !bson.IsObjectIdHex(organizationID) {
		return nil, &models.HandlerError{nil, "organizationID is not valid", http.StatusBadRequest}
	}
	err, organization := models.GetOrganization(organizationID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find organization " + organizationID, http.StatusNotFound}
	}
	return organization, nil
}

func AddUser(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	defer r.Body.Close()
	organizationID := mux.Vars(r)["organizationID"]
	
	if !bson.IsObjectIdHex(organizationID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, organization := models.AddUserToOrganization(user, organizationID)
	
	if err != nil {
		log.Println("Error 2")
		return organization, &models.HandlerError{err, "Could not add the user ", http.StatusNotFound}
	}
	log.Println("Sin error")
	return organization, nil
}