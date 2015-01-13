package controllers

import (
	models "../models"
	"encoding/json"
	"net/http"
)

func ListDesigners(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	designers, _ := models.AllDesigners()
	if designers == nil {
		return []models.Designer{}, nil
	}
	return designers, nil
}

func GetDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	// mux.Vars grabs variables from the path
	// Id := mux.Vars(r)["id"]
	Id := r.URL.Path[len("/designers/"):]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, b := models.GetDesigner(Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + Id, http.StatusNotFound}
	}
	return b, nil
}

func AddDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var payload models.Designer
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}

	err, designer := models.CreateDesigner(payload)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create designer ", http.StatusNotFound}
	}
	return designer, nil
}

func UpdateDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := r.URL.Path[len("/designers/"):]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	var payload models.Designer
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, designer := models.UpdateDesigner(payload, Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update designer " + Id + " to update", http.StatusNotFound}
	}
	return designer, nil
}

func RemoveDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := r.URL.Path[len("/designers/"):]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveDesigner(Id)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + Id + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
