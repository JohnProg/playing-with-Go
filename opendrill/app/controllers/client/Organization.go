package client

import (
	models "../../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

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

func AddUserToOrganzation(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
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
		return organization, &models.HandlerError{err, "Could not add the user ", http.StatusNotFound}
	}
	return organization, nil
}

func AddOrganization(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError){
	defer r.Body.Close()
	userID := mux.Vars(r)["organizatorID"]
	
	if !bson.IsObjectIdHex(userID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}

	var organization models.Organization
	if err := json.NewDecoder(r.Body).Decode(&organization); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}

	err, organization := models.AddOrganization(userID, organization)
	
	if err != nil {
		return organization, &models.HandlerError{err, "Could not add the organization ", http.StatusNotFound}
	}
	return organization, nil
}