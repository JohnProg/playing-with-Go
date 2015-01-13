package controllers

import (
	models "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func AllOrganizators(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizators, _ := models.AllOrganizators()
	if organizators == nil {
		return []models.Organizator{}, nil
	}
	return organizators, nil
}

func GetOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	if !bson.IsObjectIdHex(organizatorID) {
		return nil, &models.HandlerError{nil, "organizatorID is not valid", http.StatusBadRequest}
	}
	err, b := models.GetOrganizator(organizatorID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + organizatorID, http.StatusNotFound}
	}
	return b, nil
}

func CreateOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var organizator models.Organizator
	if err := json.NewDecoder(r.Body).Decode(&organizator); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}

	err, designer := models.CreateOrganizator(organizator)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create designer ", http.StatusNotFound}
	}
	return designer, nil
}

func UpdateOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	if !bson.IsObjectIdHex(organizatorID) {
		return nil, &models.HandlerError{nil, "organizatorID is not valid", http.StatusBadRequest}
	}
	var organizator models.Organizator
	if err := json.NewDecoder(r.Body).Decode(&organizator); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, designer := models.UpdateOrganizator(organizator, organizatorID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update designer " + organizatorID + " to update", http.StatusNotFound}
	}
	return designer, nil
}

func RemoveOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	if !bson.IsObjectIdHex(organizatorID) {
		return nil, &models.HandlerError{nil, "organizatorID is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveOrganizator(organizatorID)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + organizatorID + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
