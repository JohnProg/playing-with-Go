package controllers

import (
	models "../models"
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
	templates := models.GetOrganizationsFromOrganizator(organizatorID)
	if templates == nil {
		return []models.Template{}, nil
	}
	return templates, nil
}

func GetOrganizationFromOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	organizationID := mux.Vars(r)["organizationID"]
	if !bson.IsObjectIdHex(organizatorID) && !bson.IsObjectIdHex(organizationID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	err, template := models.GetOrganizationFromOrganizator(organizatorID, organizationID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + organizationID, http.StatusNotFound}
	}
	return template, nil
}

func AddOrganizationToOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	if !bson.IsObjectIdHex(organizatorID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, template := models.AddOrganizationToOrganizator(template, organizatorID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create book ", http.StatusNotFound}
	}
	return template, nil
}

func UpdateOrganizationFromOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	organizationID := mux.Vars(r)["organizationID"]
	if !bson.IsObjectIdHex(organizatorID) && !bson.IsObjectIdHex(organizationID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, template := models.UpdateOrganizationFromOrganizator(template, organizatorID, organizationID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update book " + organizatorID + " to update", http.StatusNotFound}
	}
	return template, nil
}

func RemoveOrganizationFromOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	organizationID := mux.Vars(r)["organizationID"]
	if !bson.IsObjectIdHex(organizatorID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	err, deleted := models.RemoveOrganizationFromOrganizator(organizatorID, organizationID)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find book " + organizatorID + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
