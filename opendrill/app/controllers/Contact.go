package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	models "../models"
)

func ListContacts(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	contacts, _ := models.AllContact()
	if contacts == nil {
		return []models.Contact{}, nil
	}
	return contacts, nil
}

func GetContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	// mux.Vars grabs variables from the path
	contactID := mux.Vars(r)["contactID"]
	if !bson.IsObjectIdHex(contactID) {
		return nil, &models.HandlerError{nil, "contactID is not valid", http.StatusBadRequest}
	}
	err, b := models.GetContact(contactID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find contact " + contactID, http.StatusNotFound}
	}
	return b, nil
}

func AddContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, contact := models.CreateContact(contact)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create contact ", http.StatusNotFound}
	}
	return contact, nil
}

func UpdateContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	contactID := mux.Vars(r)["contactID"]
	if !bson.IsObjectIdHex(contactID) {
		return nil, &models.HandlerError{nil, "contactID is not valid", http.StatusBadRequest}
	}
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, contact := models.UpdateContact(contact, contactID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update contact " + contactID + " to update", http.StatusNotFound}
	}
	return contact, nil
}

func RemoveContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	contactID := mux.Vars(r)["contactID"]
	if !bson.IsObjectIdHex(contactID) {
		return nil, &models.HandlerError{nil, "contactID is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveContact(contactID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find contact " + contactID + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
