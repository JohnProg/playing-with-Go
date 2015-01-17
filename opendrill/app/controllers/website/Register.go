package website

import (
	models "../../models"
	"encoding/json"
	"net/http"
)


func RegisterUser(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var user models.User
		user.Role = 0

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON", http.StatusNotFound}
	}

	err, organization := models.CreateOrganization(user)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create organization: ", http.StatusNotFound}
	}
	return organization, nil
}