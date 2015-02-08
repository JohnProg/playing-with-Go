package website

import (
	models "../../models"
	"encoding/json"
	"net/http"
)

func AuthUser(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON", http.StatusNotFound}
	}

	err, organization := models.AuthUser(user)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create organization: ", http.StatusNotFound}
	}
	return organization, nil
}
