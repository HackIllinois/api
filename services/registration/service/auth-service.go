package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"net/http"
)

/*
	Add applicant role to user with auth service
*/
func AddApplicantRole(id string) error {
	return AddRole(id, "Applicant")
}

/*
	Add mentor role to user with auth service
*/
func AddMentorRole(id string) error {
	return AddRole(id, "Mentor")
}

/*
	Add role to user with auth service
*/
func AddRole(id string, role string) error {
	user_role_modification := models.UserRoleModification{ID: id, Role: role}

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
