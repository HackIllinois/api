package service

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
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
	// Grab pagination and sorting parameters and delete to prevent the CreateFilterQuery from using them
	page := parameters["p"]
	page_limit := parameters["limit"]
	sort_parameters := parameters["sortby"]
	delete(parameters, "p")
	delete(parameters, "limit")
	delete(parameters, "sortby")

	query, err := database.CreateFilterQuery(parameters, models.UserInfo{})

	if err != nil {
		return nil, err
	}

	var filtered_users models.FilteredUsers

	if len(sort_parameters) == 1 {
		// Split by comma to get a slice of sorting parameters
		// i.e FirstName, LastName --> ["FirstName", "LastName"]
		sort_parameters = strings.Split(sort_parameters[0], ",")

		var sort_fields []database.SortField

		for _, field := range sort_parameters {
			// Push to lowercase because MongoDB columns are all lowercase
			field = strings.ToLower(field)
			field = strings.TrimSpace(field)

			sort_fields = append(sort_fields,
				database.SortField{
					Name:     field,
					Reversed: false,
				})
		}

		// Fetch and Sort
		err = db.FindAllSorted("info", query, sort_fields, &filtered_users.Users)
	} else {
		err = db.FindAll("info", query, &filtered_users.Users)
	}

	if err != nil {
		return nil, err
	}

	if len(page) == 1 && len(page_limit) == 1 {
		page, _ := strconv.Atoi(page[0])
		page_limit, _ := strconv.Atoi(page_limit[0])

		// Subtract one because page numbers will be 1-indexed
		// The max() function will ensure we don't paginate past the length of the Users list
		filtered_users.Users = filtered_users.Users[(page-1)*page_limit : utils.Max(page*page_limit, len(filtered_users.Users))]
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
