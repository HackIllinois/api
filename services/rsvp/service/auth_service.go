package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/models"
	"net/http"
)

/*
	Add Attendee role to user with auth service
*/
func AddAttendeeRole(id string) error {
	var user_roles models.UserRoles
	status, err := apirequest.Get(config.AUTH_SERVICE+"/auth/roles/"+id+"/", &user_roles)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	user_roles.Roles = append(user_roles.Roles, "Attendee")

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_roles)

	status, err = apirequest.Put(config.AUTH_SERVICE+"/auth/roles/", &body, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	return nil
}

/*
	Remove Attendee role from user with auth service
*/
func RemoveAttendeeRole(id string) error {
	var user_roles models.UserRoles
	status, err := apirequest.Get(config.AUTH_SERVICE+"/auth/roles/"+id+"/", &user_roles)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	for index, role := range user_roles.Roles {
		if role == "Attendee" {
			user_roles.Roles = append(user_roles.Roles[:index], user_roles.Roles[index+1:]...)
		}
	}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_roles)

	status, err = apirequest.Put(config.AUTH_SERVICE+"/auth/roles/", &body, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	return nil
}
