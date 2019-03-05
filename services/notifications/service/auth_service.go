package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"net/http"
)

/*
	Gets the list of valid roles
*/
func GetValidRoles() (*models.UserRoleList, error) {
	var user_roles_list models.UserRoleList
	status, err := apirequest.Get(config.AUTH_SERVICE+"/auth/roles/list/", &user_roles_list)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Could not fetch list of valid roles from Auth service.")
	}

	return &user_roles_list, nil
}

/*
	Gets the list of valid roles
*/
func GetUsersByRole(role string) ([]string, error) {
	var user_list models.UserList
	status, err := apirequest.Get(config.AUTH_SERVICE+"/auth/roles/list/"+role+"/", &user_list)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Could not fetch list of user by role from Auth service.")
	}

	return user_list.UserIDs, nil
}

/*
	Gets the roles for a given user
*/
func GetUserRoles(id string) ([]string, error) {
	var user_roles models.UserRoles
	status, err := apirequest.Get(config.AUTH_SERVICE+"/auth/roles/"+id+"/", &user_roles)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Could not fetch user's roles from Auth service.")
	}

	return user_roles.Roles, nil
}
