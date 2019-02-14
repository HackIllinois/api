package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/models"
	"net/http"
)

/*
	Gets the roles for a user given id.
*/
func GetRoles(id string) (*models.UserRoles, error) {
	var user_roles models.UserRoles
	status, err := apirequest.Get(config.AUTH_SERVICE+"/auth/roles/"+id+"/", &user_roles)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Could not fetch roles from Auth service.")
	}

	return &user_roles, nil
}
