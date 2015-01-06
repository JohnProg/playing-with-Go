package controllers

import (
	"net/http"
	"encoding/json"

	models "../models"

)


func ListListContacts(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	list_contacts, _ := models.AllListContact()
	if list_contacts == nil {
		return []models.ListContact{}, nil
	}
	return list_contacts, nil
}

func AddListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var payload models.ListContact

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON", http.StatusNotFound}
	}
	err, list_contact := models.CreateListContact(payload)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create list contact", http.StatusNotFound}
	}
	return list_contact, nil
} 