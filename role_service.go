package main

import (
	"github.com/hackillinois/api-auth/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	Get the user's roles by id
	If the user has no roles and create_user is true they will be assigned the role User
	This generally occurs the first time the user logs into the service
*/
func GetUserRoles(id string, create_user bool) ([]string, error) {
	query := bson.M {
		"id": id,
	}

	var roles models.UserRoles
	err := FindOne("roles", query, &roles)

	if err != nil {
		if err == mgo.ErrNotFound && create_user {
			Insert("roles", &models.UserRoles {
				ID: id,
				Roles: []string{"User"},
			})

			err := FindOne("roles", query, &roles)

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
	Sets the roles for the user with the specified id
*/
func SetUserRoles(id string, roles []string) error {
	selector := bson.M {
		"id": id,
	}

	err := Update("roles", selector, &models.UserRoles {
		ID: id,
		Roles: roles,
	})

	return err
}
