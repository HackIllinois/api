package service

import (
	"errors"
	"strings"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
)

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.AUTH_DB_HOST, config.AUTH_DB_NAME)

	if err != nil {
		return err
	}

	return nil
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
	err := db.FindOne("roles", query, &roles, nil)

	if err != nil {
		if err == database.ErrNotFound && create_user {
			db.Insert("roles", &models.UserRoles{
				ID:    id,
				Roles: []string{"User"},
			}, nil)

			err := db.FindOne("roles", query, &roles, nil)

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

	if !utils.ContainsString(roles, role) {
		roles = append(roles, role)
	}

	err = db.Replace("roles", selector, &models.UserRoles{
		ID:    id,
		Roles: roles,
	}, false, nil)

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

	roles, err = utils.RemoveString(roles, role)

	if err != nil {
		return errors.New("User does not have specified role")
	}

	err = db.Replace("roles", selector, &models.UserRoles{
		ID:    id,
		Roles: roles,
	}, false, nil)

	return err
}

/*
	Automatically grant staff and admin roles based on user's verified email
*/
func AddAutomaticRoleGrants(id string, email string) error {
	email_components := strings.Split(email, "@")

	if len(email_components) < 2 {
		return errors.New("Could not parse user's email")
	}

	domain := email_components[1]

	if domain == config.STAFF_DOMAIN {
		err := AddUserRole(id, models.StaffRole)

		if err != nil {
			return err
		}
	}

	if email == config.SYSTEM_ADMIN_EMAIL {
		err := AddUserRole(id, models.AdminRole)

		if err != nil {
			return err
		}
	}

	return nil
}

/*
	Returns a list of valid roles for a user to be assigned
*/
func GetValidRoles() []models.Role {
	return models.Roles
}

/*
	Returns a list of user ids with a given role
*/
func GetUsersByRole(role models.Role) ([]string, error) {
	query := database.QuerySelector{
		"roles": database.QuerySelector{
			"$elemMatch": database.QuerySelector{
				"$eq": role,
			},
		},
	}

	var users []models.UserRoles
	err := db.FindAll("roles", query, &users, nil)

	if err != nil {
		return nil, err
	}

	userids := make([]string, len(users))

	for i, user := range users {
		userids[i] = user.ID
	}

	return userids, nil
}

/*
	Returns role stats
*/
func GetStats() (map[string]interface{}, error) {
	stats, err := db.GetStats("roles", []string{"roles"}, nil)
	if err != nil {
		return nil, err
	}
	return stats, nil
}
