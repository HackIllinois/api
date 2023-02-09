package main

import (
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/utilities/accountgen/models"
)

var auth_db database.Database
var user_db database.Database

func init() {
	auth_db_connection, err := database.InitDatabase("mongodb://localhost", "auth")

	if err != nil {
		panic(err)
	}

	user_db_connection, err := database.InitDatabase("mongodb://localhost", "user")

	if err != nil {
		panic(err)
	}

	auth_db = auth_db_connection
	user_db = user_db_connection
}

func CreateAccount(id string, roles []string, username string, firstName string, lastName string, email string) error {
	err := PopulateAuthInfo(id, roles)

	if err != nil {
		return err
	}

	err = PopulateUserInfo(id, username, firstName, lastName, email)

	return err
}

func PopulateAuthInfo(id string, roles []string) error {
	user_roles := models.UserRoles{
		ID:    id,
		Roles: roles,
	}

	selector := database.QuerySelector{
		"id": id,
	}

	err := auth_db.Replace("roles", selector, &user_roles, true, nil)

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

	selector := database.QuerySelector{
		"id": id,
	}

	err := user_db.Replace("info", selector, &user_info, true, nil)

	return err
}
