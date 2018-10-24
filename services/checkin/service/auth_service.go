package service

import (
	"bytes"
	"encoding/json"
	"errors"
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

	client := http.Client{}
	req, err := http.NewRequest("PUT", config.AUTH_SERVICE+"/auth/roles/add/", &body)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	return nil
}

/*
	Remove Attendee role from user with auth service
*/
func RemoveAttendeeRole(id string) error {
	user_role_modification := models.UserRoleModification{ID: id, Role: "Attendee"}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_role_modification)

	client := http.Client{}
	req, err := http.NewRequest("PUT", config.AUTH_SERVICE+"/auth/roles/remove/", &body)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	return nil
}
