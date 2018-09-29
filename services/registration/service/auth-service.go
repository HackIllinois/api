package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ethan-lord/api/services/registration/config"
	"github.com/ethan-lord/api/services/registration/models"
	"net/http"
)

/*
	Add applicant role to user with auth service
*/
func AddApplicantRole(id string) error {
	return AddRole(id, models.Applicant)
}

/*
	Add mentor role to user with auth service
*/
func AddMentorRole(id string) error {
	return AddRole(id, models.Mentor)
}

/*
	Add role to user with auth service
*/
func AddRole(id string, role models.Role) error {
	resp, err := http.Get(config.AUTH_SERVICE + "/auth/roles/" + id + "/")

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	var user_roles models.UserRoles
	json.NewDecoder(resp.Body).Decode(&user_roles)

	user_roles.Roles = append(user_roles.Roles, role)

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
