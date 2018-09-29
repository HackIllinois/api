package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ethan-lord/api/services/rsvp/config"
	"github.com/ethan-lord/api/services/rsvp/models"
	"net/http"
)

/*
	Add Attendee role to user with auth service
*/
func AddAttendeeRole(id string) error {
	resp, err := http.Get(config.AUTH_SERVICE + "/auth/roles/" + id + "/")

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	var user_roles models.UserRoles
	json.NewDecoder(resp.Body).Decode(&user_roles)

	user_roles.Roles = append(user_roles.Roles, "Attendee")

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_roles)

	client := http.Client{}
	req, err := http.NewRequest("PUT", config.AUTH_SERVICE+"/auth/roles/", &body)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)

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
	resp, err := http.Get(config.AUTH_SERVICE + "/auth/roles/" + id + "/")

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	var user_roles models.UserRoles
	json.NewDecoder(resp.Body).Decode(&user_roles)

	for index, role := range user_roles.Roles {
		if role == "Attendee" {
			user_roles.Roles = append(user_roles.Roles[:index], user_roles.Roles[index+1:]...)
		}
	}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_roles)

	client := http.Client{}
	req, err := http.NewRequest("PUT", config.AUTH_SERVICE+"/auth/roles/", &body)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	return nil
}
