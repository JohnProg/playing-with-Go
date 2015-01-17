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
	organizations := models.GetOrganizationsFromOrganizator(organizatorID)
	if organizations == nil {
		return []models.Organization{}, nil
	}
	return organizations, nil
}

func GetOrganizationFromOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	organizationID := mux.Vars(r)["organizationID"]
	if !bson.IsObjectIdHex(organizatorID) && !bson.IsObjectIdHex(organizationID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	err, organization := models.GetOrganizationFromOrganizator(organizatorID, organizationID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + organizationID, http.StatusNotFound}
	}
	return organization, nil
}

func AddOrganizationToOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	if !bson.IsObjectIdHex(organizatorID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	defer r.Body.Close()
	var organization models.Organization
	if err := json.NewDecoder(r.Body).Decode(&organization); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, organization := models.AddOrganizationToOrganizator(organization, organizatorID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create book ", http.StatusNotFound}
	}
	return organization, nil
}

func UpdateOrganizationFromOrganizator(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	organizatorID := mux.Vars(r)["organizatorID"]
	organizationID := mux.Vars(r)["organizationID"]
	if !bson.IsObjectIdHex(organizatorID) && !bson.IsObjectIdHex(organizationID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	defer r.Body.Close()
	var organization models.Organization
	if err := json.NewDecoder(r.Body).Decode(&organization); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, organization := models.UpdateOrganizationFromOrganizator(organization, organizatorID, organizationID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update book " + organizatorID + " to update", http.StatusNotFound}
	}
	return organization, nil
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
