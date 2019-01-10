package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"net/http"
)

/*
	Add applicant role to user with auth service
*/
func AddApplicantRole(id string) error {
	return AddRole(id, models.ApplicantRole)
}

/*
	Add mentor role to user with auth service
*/
func AddMentorRole(id string) error {
	return AddRole(id, models.MentorRole)
}

/*
	Add role to user with auth service
*/
func AddRole(id string, role models.Role) error {
	user_role_modification := models.UserRoleModification{ID: id, Role: role}

	status, err := apirequest.Put(config.AUTH_SERVICE+"/auth/roles/add/", &user_role_modification, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Auth service failed to update roles.")
	}

	return nil
}
