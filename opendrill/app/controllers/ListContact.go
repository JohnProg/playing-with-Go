package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	models "../models"
)

func ListListContacts(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	list_contacts, _ := models.AllListContact()
	if list_contacts == nil {
		return []models.ListContact{}, nil
	}
	return list_contacts, nil
}

func GetListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	listContactID := mux.Vars(r)["listContactID"]
	if !bson.IsObjectIdHex(listContactID) {
		return nil, &models.HandlerError{nil, "listContactID es not valid", http.StatusBadRequest}
	}
	err, b := models.GetListContact(listContactID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find list contact" + listContactID, http.StatusNotFound}
	}
	return b, nil
}

func AddListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var contactList models.ListContact

	if err := json.NewDecoder(r.Body).Decode(&contactList); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON", http.StatusNotFound}
	}
	err, list_contact := models.CreateListContact(contactList)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create list contact", http.StatusNotFound}
	}
	return list_contact, nil
}

func UpdateListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	listContactID := mux.Vars(r)["listContactID"]
	if !bson.IsObjectIdHex(listContactID) {
		return nil, &models.HandlerError{nil, "listContactID is not valid", http.StatusBadRequest}
	}
	var contactList models.ListContact
	if err := json.NewDecoder(r.Body).Decode(&contactList); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON", http.StatusNotFound}
	}
	err, list_contact := models.UpdateListContact(contactList, listContactID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update list contact" + listContactID + "to update", http.StatusNotFound}
	}
	return list_contact, nil
}

func RemoveListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	listContactID := mux.Vars(r)["listContactID"]
	if !bson.IsObjectIdHex(listContactID) {
		return nil, &models.HandlerError{nil, "listContactID is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveListContact(listContactID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find list contact" + listContactID + " to delete", http.StatusNotFound}
	}
	return deleted, nil

}
