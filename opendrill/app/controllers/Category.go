package controllers

import (
<<<<<<< HEAD
	"github.com/gorilla/mux"

	"net/http"
	"log"
	"encoding/json"
=======
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada

	models "../models"
)

<<<<<<< HEAD

=======
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada
func ListCategories(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	categories, _ := models.AllCategory()
	if categories == nil {
		return []models.Category{}, nil
	}
	return categories, nil
}

func GetCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	// mux.Vars grabs variables from the path
<<<<<<< HEAD
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, b := models.GetCategory(Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find category " + Id, http.StatusNotFound}
=======
	categoryID := mux.Vars(r)["categoryID"]
	if !bson.IsObjectIdHex(categoryID) {
		return nil, &models.HandlerError{nil, "categoryID is not valid", http.StatusBadRequest}
	}
	err, b := models.GetCategory(categoryID)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find category " + categoryID, http.StatusNotFound}
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada
	}
	return b, nil
}

func AddCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
<<<<<<< HEAD
	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, category := models.CreateCategory(payload)
=======
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, category := models.CreateCategory(category)
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create category ", http.StatusNotFound}
	}
	return category, nil
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
<<<<<<< HEAD
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, category := models.UpdateCategory(payload, Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update category " + Id + " to update", http.StatusNotFound}
=======
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
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada
	}
	return category, nil
}

func RemoveCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
<<<<<<< HEAD
	Id := mux.Vars(r)["id"]
	log.Println(Id)
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveCategory(Id)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find category " + Id + " to delete", http.StatusNotFound}
=======
	categoryID := mux.Vars(r)["categoryID"]
	if !bson.IsObjectIdHex(categoryID) {
		return nil, &models.HandlerError{nil, "categoryID is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveCategory(categoryID)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find category " + categoryID + " to delete", http.StatusNotFound}
>>>>>>> e08185a76af3c54738ab1eabc6600135d2d7dada
	}
	return deleted, nil
}
