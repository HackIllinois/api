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
	user_role_modification := models.UserRoleModification{ID: id, Role: "Attendee"}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_role_modification)

	status, err := apirequest.Put(config.AUTH_SERVICE+"/auth/roles/add/", &body, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	return nil
}
