package service

import (
	"errors"
	"net/url"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/models"
)

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.USER_DB_HOST, config.USER_DB_NAME)

	if err != nil {
		return err
	}

	return nil
}

/*
	Returns the info associated with the given user id
*/
func GetUserInfo(id string) (*models.UserInfo, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var user_info models.UserInfo
	err := db.FindOne("info", query, &user_info)

	if err != nil {
		return nil, err
	}

	return &user_info, nil
}

/*
	Set the info associated with the given user id
	The record will be created if it does not already exist
*/
func SetUserInfo(id string, user_info models.UserInfo) error {
	selector := database.QuerySelector{
		"id": id,
	}

	err := db.Update("info", selector, &user_info)

	if err == database.ErrNotFound {
		err = db.Insert("info", &user_info)
	}

	return err
}

/*
	Returns the users associated with the given parameters
*/
func GetFilteredUserInfo(parameters map[string][]string) (*models.FilteredUsers, error) {
	query, err := database.CreateFilterQuery(parameters, models.UserInfo{})

	if err != nil {
		return nil, err
	}

	var filtered_users models.FilteredUsers
	err = db.FindAll("info", query, &filtered_users.Users)

	if err != nil {
		return nil, err
	}

	return &filtered_users, nil
}

/*
	Generates a QR string for a user with the provided ID, as a URI
*/
func GetQrInfo(id string) (string, error) {
	_, err := GetUserInfo(id)

	if err != nil {
		return "", errors.New("User does not exist.")
	}

	// Construct the URI

	uri, err := url.Parse("hackillinois://user")

	if err != nil {
		return "", err
	}

	// All the fields that will be embedded in the QR code URI
	parameters := url.Values{
		"userId": []string{id},
	}

	uri.RawQuery = parameters.Encode()

	return uri.String(), nil
}

/*
	Returns all user stats
*/
func GetStats() (map[string]interface{}, error) {
	return db.GetStats("info", []string{})
}
