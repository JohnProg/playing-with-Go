package client

import (
	models "../../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

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
	user.Id = bson.NewObjectId()
	err, organization := models.AddUserToOrganization(user, organizationID)
	
	if err != nil {
		return organization, &models.HandlerError{err, "Could not add the user ", http.StatusNotFound}
	}
	return organization, nil
}

func DeleteUserToOrganzation(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	defer r.Body.Close()
	organizationID := mux.Vars(r)["organizationID"]
	userID := mux.Vars(r)["userID"]
	
	if !bson.IsObjectIdHex(organizationID) && !bson.IsObjectIdHex(userID){
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}

	err, deleted := models.DeleteUserToOrganization(userID, organizationID)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not delete user " + userID + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}

