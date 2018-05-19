package service

import (
	"errors"
	"github.com/HackIllinois/api-rsvp/database"
	"github.com/HackIllinois/api-rsvp/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	Returns the rsvp associated with the given user id
*/
func GetUserRsvp(id string) (*models.UserRsvp, error) {
	query := bson.M{
		"id": id,
	}

	var rsvp models.UserRsvp
	err := database.FindOne("rsvps", query, &rsvp)

	if err != nil {
		return nil, err
	}

	return &rsvp, nil
}

/*
	Creates the rsvp associated with the given user id
*/
func CreateUserRsvp(id string, rsvp models.UserRsvp) error {
	_, err := GetUserRsvp(id)

	if err != mgo.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Rsvp already exists")
	}

	err = database.Insert("rsvps", &rsvp)

	return err
}

/*
	Updates the rsvp associated with the given user id
*/
func UpdateUserRsvp(id string, rsvp models.UserRsvp) error {
	selector := bson.M{
		"id": id,
	}

	err := database.Update("rsvps", selector, &rsvp)

	return err
}
