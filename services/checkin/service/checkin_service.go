package service

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/models"
)

var db database.Database

func init() {
	db_connection, err := database.InitDatabase(config.CHECKIN_DB_HOST, config.CHECKIN_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns the checkin associated with the given user id
*/
func GetUserCheckin(id string) (*models.UserCheckin, error) {
	query := database.QuerySelector{
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

	if err != database.ErrNotFound {
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
	selector := database.QuerySelector{
		"id": id,
	}

	err := db.Update("checkins", selector, &user_checkin)

	return err
}

/*
	Generates a QR string for a user with the provided ID, as a URI
*/
func GetQrInfo(id string) (string, error) {

	// Retrieve all the info that needs to be embedded

	checkin_status, err := GetUserCheckin(id)

	if err != nil {
		return "", err
	}

	// Construct the URI

	uri, err := url.Parse("hackillinois://info")

	if err != nil {
		return "", err
	}

	// All the fields that will be embedded in the QR code URI
	parameters := url.Values{
		"userId":          []string{id},
		"hasCheckedIn":    []string{strconv.FormatBool(checkin_status.HasCheckedIn)},
		"hasPickedUpSwag": []string{strconv.FormatBool(checkin_status.HasPickedUpSwag)},
	}

	uri.RawQuery = parameters.Encode()

	return uri.String(), nil
}

/*
	Returns true, nil if a user with specified ID is allowed to checkin, and false, nil if not allowed.
*/
func CanUserCheckin(id string, user_has_override bool) (bool, error) {
	is_user_registered, err := IsUserRegistered(id)

	if err != nil {
		return false, err
	}

	// To checkin, the user must either (have RSVPed) or (have registered and got an override)
	if is_user_registered && user_has_override {
		return true, nil
	}

	// We do not want to call the below service function if the above condition is met, as it results
	// in a 400 (Bad Request) / error if the user's RSVP info cannot be found.
	// Therefore, we do not combine the conditions, and return as early as possible.
	is_user_rsvped, err := IsAttendeeRsvped(id)

	if err != nil {
		return false, err
	}

	return is_user_rsvped, nil
}
