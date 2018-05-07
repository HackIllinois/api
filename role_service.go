package main

import (
	"./models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	Get the user's roles by id
	If the user has no roles they will be assigned the role User
	This occurs the first time the user logs into the service
*/
func GetUserRoles(id string) ([]string, error) {
	query := bson.M {
		"id": id,
	}

	var roles models.UserRoles
	err := FindOne("roles", query, &roles)

	if err != nil {
		if err == mgo.ErrNotFound {
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
