package client

import (
	models "../../models"
	"log"
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
	err, organizations := models.GetOrganizationsFromOrganizator(organizatorID)
	if err != nil {
		return []models.Organization{}, nil
	}
	return organizations, nil
}