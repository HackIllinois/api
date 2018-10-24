package service

import (
	"errors"
	"github.com/HackIllinois/api/common/utils"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
)

var db database.Database

func init() {
	db_connection, err := database.InitDatabase(config.AUTH_DB_HOST, config.AUTH_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Get the user's roles by id
	If the user has no roles and create_user is true they will be assigned the role User
	This generally occurs the first time the user logs into the service
*/
func GetUserRoles(id string, create_user bool) ([]string, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var roles models.UserRoles
	err := db.FindOne("roles", query, &roles)

	if err != nil {
		if err == database.ErrNotFound && create_user {
			db.Insert("roles", &models.UserRoles{
				ID:    id,
				Roles: []string{"User"},
			})

			err := db.FindOne("roles", query, &roles)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return roles.Roles, nil
}

/*
	Adds a role to the user with the specified id
*/
func AddUserRole(id string, role string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	roles, err := GetUserRoles(id, false)

	if err != nil {
		return err
	}

	if !slice_utils.ContainsString(roles, role) {
		roles = append(roles, role)
	}

	err = db.Update("roles", selector, &models.UserRoles{
		ID:    id,
		Roles: roles,
	})

	return err
}

/*
	Removes a role from the user with the specified id
*/
func RemoveUserRole(id string, role string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	roles, err := GetUserRoles(id, false)

	if err != nil {
		return err
	}

	roles, err = slice_utils.RemoveString(roles, role)

	if err != nil {
		return errors.New("User does not have specified role")
	}

	err = db.Update("roles", selector, &models.UserRoles{
		ID:    id,
		Roles: roles,
	})

	return err
}
