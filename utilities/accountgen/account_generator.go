package main

import (
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/utilities/accountgen/models"
	"gopkg.in/mgo.v2/bson"
)

var auth_db database.MongoDatabase
var user_db database.MongoDatabase

func init() {
	auth_db_connection, err := database.InitMongoDatabase("localhost", "auth")

	if err != nil {
		panic(err)
	}

	user_db_connection, err := database.InitMongoDatabase("localhost", "user")

	if err != nil {
		panic(err)
	}

	auth_db = auth_db_connection
	user_db = user_db_connection
}

func CreateAccount(id string, roles []models.Role, username string, firstName string, lastName string, email string) error {
	err := PopulateAuthInfo(id, roles)

	if err != nil {
		return err
	}

	err = PopulateUserInfo(id, username, firstName, lastName, email)

	return err
}

func PopulateAuthInfo(id string, roles []models.Role) error {
	user_roles := models.UserRoles{
		ID:    id,
		Roles: roles,
	}

	selector := bson.M{
		"id": id,
	}

	_, err := auth_db.Upsert("roles", selector, &user_roles)

	return err
}

func PopulateUserInfo(id string, username string, firstName string, lastName string, email string) error {
	user_info := models.UserInfo{
		ID:        id,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	selector := bson.M{
		"id": id,
	}

	_, err := user_db.Upsert("info", selector, &user_info)

	return err
}
