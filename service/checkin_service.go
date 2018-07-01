package service

import (
	"errors"
	"strconv"
	"net/url"
	"github.com/HackIllinois/api-checkin/config"
	"github.com/HackIllinois/api-checkin/models"
	"github.com/HackIllinois/api-commons/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.CHECKIN_DB_HOST, config.CHECKIN_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns the checkin associated with the given user id
*/
func GetUserCheckin(id string) (*models.UserCheckin, error) {
	query := bson.M{
		"id": id,
	}

	var user_checkin models.UserCheckin
	err := db.FindOne("checkins", query, &user_checkin)

	if err != nil {
		return nil, err
	}

	return &user_checkin, nil
}

/*
	Create the checkin associated with the given user id
*/
func CreateUserCheckin(id string, user_checkin models.UserCheckin) error {
	_, err := GetUserCheckin(id)

	if err != mgo.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Checkin already exists")
	}

	err = db.Insert("checkins", &user_checkin)

	return err
}

/*
	Update the checkin associated with the given user id
*/
func UpdateUserCheckin(id string, user_checkin models.UserCheckin) error {
	selector := bson.M{
		"id": id,
	}

	err := db.Update("checkins", selector, &user_checkin)

	return err
}

/*
	Generates a QR string for a user with the provided ID, as a URI
*/
func GetQrString(id string) (string, error) {

	// Retrieve all the info that needs to be embedded
	
	checkin_status, err := GetUserCheckin(id)

	if err != nil {
		return "", err
	}

	has_checked_in_string := strconv.FormatBool(checkin_status.HasCheckedIn)

	has_picked_up_swag_string := strconv.FormatBool(checkin_status.HasPickedUpSwag)

	// Construct the URI

	uri, err := url.Parse("hackillinois://info")

	if err != nil {
		return "", err
	}

	// All the fields that will be embedded in the QR code URI
	parameters := url.Values{
		"userId": []string{id},
		"hasCheckedIn": []string{has_checked_in_string},
		"hasPickedUpSwag": []string{has_picked_up_swag_string},
	}

	uri.RawQuery = parameters.Encode()

	return uri.String(), nil
}
