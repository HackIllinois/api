package service

import (
	"errors"
	"github.com/HackIllinois/api-registration/database"
	"github.com/HackIllinois/api-registration/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	Returns the registration associated with the given user id
*/
func GetUserRegistration(id string) (*models.UserRegistration, error) {
	query := bson.M{
		"id": id,
	}

	var user_registration models.UserRegistration
	err := database.FindOne("attendees", query, &user_registration)

	if err != nil {
		return nil, err
	}

	return &user_registration, nil
}

/*
	Creates the registration associated with the given user id
*/
func CreateUserRegistration(id string, user_registration models.UserRegistration) error {
	_, err := GetUserRegistration(id)

	if err != mgo.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Registration already exists")
	}

	err = database.Insert("attendees", &user_registration)

	return err
}

/*
	Updates the registration associated with the given user id
*/
func UpdateUserRegistration(id string, user_registration models.UserRegistration) error {
	selector := bson.M{
		"id": id,
	}

	err := database.Update("attendees", selector, &user_registration)

	return err
}
