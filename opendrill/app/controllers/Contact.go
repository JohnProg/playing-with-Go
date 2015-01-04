package controllers

import (
	models "../models"
	"net/http"
	"fmt"
)


func ListContacts(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	fmt.Println("llego")
	contacts, _ := models.AllContact()
	if contacts == nil {
		return []models.Contact{}, nil
	}
	return contacts, nil
}