package controllers

import (
	models "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func GetTemplatesFromDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	designerID := mux.Vars(r)["designerID"]
	if !bson.IsObjectIdHex(designerID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	templates := models.GetTemplatesFromDesigner(designerID)
	if templates == nil {
		return []models.Template{}, nil
	}
	return templates, nil
}

func GetTemplateFromDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	designerID := mux.Vars(r)["designerID"]
	templateID := mux.Vars(r)["templateID"]
	if !bson.IsObjectIdHex(designerID) && !bson.IsObjectIdHex(templateID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	err, template := models.GetTemplateFromDesigner(designerID, templateID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + templateID, http.StatusNotFound}
	}
	return template, nil
}

func AddTemplateToDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	designerID := mux.Vars(r)["designerID"]
	if !bson.IsObjectIdHex(designerID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	defer r.Body.Close()
	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, template := models.AddTemplateToDesigner(template, designerID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create book ", http.StatusNotFound}
	}
	return template, nil
}

func UpdateTemplateFromDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	designerID := mux.Vars(r)["designerID"]
	templateID := mux.Vars(r)["templateID"]
	if !bson.IsObjectIdHex(designerID) && !bson.IsObjectIdHex(templateID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	defer r.Body.Close()
	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, template := models.UpdateTemplateFromDesigner(template, designerID, templateID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update book " + designerID + " to update", http.StatusNotFound}
	}
	return template, nil
}

func RemoveTemplateFromDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	designerID := mux.Vars(r)["designerID"]
	templateID := mux.Vars(r)["templateID"]
	if !bson.IsObjectIdHex(designerID) {
		return nil, &models.HandlerError{nil, "Could not parse JSON", http.StatusNotFound}
	}
	err, deleted := models.RemoveTemplateFromDesigner(designerID, templateID)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find book " + designerID + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
