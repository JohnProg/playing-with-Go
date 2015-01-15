package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	models "../models"
)

func ListCategories(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	categories, _ := models.AllCategory()
	if categories == nil {
		return []models.Category{}, nil
	}
	return categories, nil
}

func GetCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	// mux.Vars grabs variables from the path
	categoryID := mux.Vars(r)["categoryID"]
	if !bson.IsObjectIdHex(categoryID) {
		return nil, &models.HandlerError{nil, "categoryID is not valid", http.StatusBadRequest}
	}
	err, b := models.GetCategory(categoryID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find category " + categoryID, http.StatusNotFound}
	}
	return b, nil
}

func AddCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, category := models.CreateCategory(category)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create category ", http.StatusNotFound}
	}
	return category, nil
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	categoryID := mux.Vars(r)["categoryID"]
	if !bson.IsObjectIdHex(categoryID) {
		return nil, &models.HandlerError{nil, "categoryID is not valid", http.StatusBadRequest}
	}
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, category := models.UpdateCategory(category, categoryID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update category " + categoryID + " to update", http.StatusNotFound}
	}
	return category, nil
}

func RemoveCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	categoryID := mux.Vars(r)["categoryID"]
	if !bson.IsObjectIdHex(categoryID) {
		return nil, &models.HandlerError{nil, "categoryID is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveCategory(categoryID)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find category " + categoryID + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
