package service

import (
	"github.com/ReflectionsProjections/api/common/database"
	"github.com/ReflectionsProjections/api/services/auth/config"
	"github.com/ReflectionsProjections/api/services/auth/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.AUTH_DB_HOST, config.AUTH_DB_NAME)

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
	query := bson.M{
		"id": id,
	}

	var roles models.UserRoles
	err := db.FindOne("roles", query, &roles)

	if err != nil {
		if err == mgo.ErrNotFound && create_user {
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
	Sets the roles for the user with the specified id
*/
func SetUserRoles(id string, roles []string) error {
	selector := bson.M{
		"id": id,
	}

	err := db.Update("roles", selector, &models.UserRoles{
		ID:    id,
		Roles: roles,
	})

	return err
}
