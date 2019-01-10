package service

import (
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
	user_role_modification := models.UserRoleModification{ID: id, Role: models.AttendeeRole}

	status, err := apirequest.Put(config.AUTH_SERVICE+"/auth/roles/add/", &user_role_modification, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles.")
	}

	return nil
}

/*
	Remove Attendee role from user with auth service
*/
func RemoveAttendeeRole(id string) error {
	user_role_modification := models.UserRoleModification{ID: id, Role: models.AttendeeRole}

	status, err := apirequest.Put(config.AUTH_SERVICE+"/auth/roles/remove/", &user_role_modification, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	return nil
}
