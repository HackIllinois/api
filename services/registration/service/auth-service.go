package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"github.com/HackIllinois/api/common/apirequest"
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
	var user_roles models.UserRoles
	status, err := apirequest.Get(config.AUTH_SERVICE + "/auth/roles/" + id + "/", &user_roles)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	user_roles.Roles = append(user_roles.Roles, role)

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_roles)

	status, err = apirequest.Put(config.AUTH_SERVICE + "/auth/roles/", &body, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles")
	}

	return nil
}
