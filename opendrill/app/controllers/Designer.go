package controllers

import (
	models "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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
	designerID := mux.Vars(r)["designerID"]
	if !bson.IsObjectIdHex(designerID) {
		return nil, &models.HandlerError{nil, "designerID is not valid", http.StatusBadRequest}
	}
	err, b := models.GetDesigner(designerID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + designerID, http.StatusNotFound}
	}
	return b, nil
}

func AddDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	defer r.Body.Close()
	var designer models.Designer
	if err := json.NewDecoder(r.Body).Decode(&designer); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}

	err, designer := models.CreateDesigner(designer)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create designer ", http.StatusNotFound}
	}
	return designer, nil
}

func UpdateDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	designerID := mux.Vars(r)["designerID"]
	if !bson.IsObjectIdHex(designerID) {
		return nil, &models.HandlerError{nil, "designerID is not valid", http.StatusBadRequest}
	}
	defer r.Body.Close()
	var designer models.Designer
	if err := json.NewDecoder(r.Body).Decode(&designer); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, designer := models.UpdateDesigner(designer, designerID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update designer " + designerID + " to update", http.StatusNotFound}
	}
	return designer, nil
}

func RemoveDesigner(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	designerID := mux.Vars(r)["designerID"]
	if !bson.IsObjectIdHex(designerID) {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveDesigner(designerID)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find designer " + designerID + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
