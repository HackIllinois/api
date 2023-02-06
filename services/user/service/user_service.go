package service

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"

	common_config "github.com/HackIllinois/api/common/config"
	"github.com/HackIllinois/api/common/database"
	hack_errors "github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/utils"
	user_config "github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/models"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(user_config.USER_DB_HOST, user_config.USER_DB_NAME)

	if err != nil {
		return err
	}

	return nil
}

/*
Returns the info associated with the given user id
*/
func GetUserInfo(id string, sessCtx *mongo.SessionContext) (*models.UserInfo, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var user_info models.UserInfo
	err := db.FindOne("info", query, &user_info, sessCtx)

	if err != nil {
		return nil, err
	}

	return &user_info, nil
}

/*
Set the info associated with the given user id
The record will be created if it does not already exist
*/
func SetUserInfo(id string, user_info models.UserInfo, sessCtx *mongo.SessionContext) error {
	selector := database.QuerySelector{
		"id": id,
	}

	err := db.Replace("info", selector, user_info, true, sessCtx)

	return err
}

func UpsertUserInfo(id string, user_info models.UserInfo) (*models.UserInfo, *hack_errors.ApiError) {
	sess, err := db.StartSession()

	if err != nil {
		hack_err := hack_errors.InternalError(err.Error(), "Could not fetch user info by ID.")
		return nil, &hack_err
	}

	data, err := (*sess).WithTransaction(context.Background(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		err := SetUserInfo(user_info.ID, user_info, &sessCtx)

		if err != nil {
			hack_err := hack_errors.DatabaseError(err.Error(), "Could not upsert user info.")
			return nil, &hack_err
		}

		updated_info, err := GetUserInfo(user_info.ID, &sessCtx)

		if err != nil {
			hack_err := hack_errors.DatabaseError(err.Error(), "Could not fetch user info by ID.")
			return nil, &hack_err
		}

		return updated_info, nil
	})

	if err != nil {
		return nil, err.(*hack_errors.ApiError)
	}

	return data.(*models.UserInfo), nil
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

		var sort_fields bson.D

		for _, field := range sort_parameters {
			// Push to lowercase because MongoDB columns are all lowercase
			field = strings.ToLower(field)
			field = strings.TrimSpace(field)

			if len(field) > 0 {
				order := 1
				if field[0] == '-' {
					order = -1
					field = field[1:]
				}
				sort_fields = append(sort_fields, bson.E{field, order})
			}
		}

		// Fetch and Sort
		err = db.FindAllSorted("info", query, sort_fields, &filtered_users.Users, nil)
	} else {
		err = db.FindAll("info", query, &filtered_users.Users, nil)
	}

	if err != nil {
		return nil, err
	}

	if len(page) == 1 && len(page_limit) == 1 {
		page, _ := strconv.Atoi(page[0])
		page_limit, _ := strconv.Atoi(page_limit[0])

		// Subtract one because page numbers will be 1-indexed
		// The max() function will ensure we don't paginate past the length of the Users list
		filtered_users.Users = filtered_users.Users[(page-1)*page_limit : utils.Min(page*page_limit, len(filtered_users.Users))]
	}

	return &filtered_users, nil
}

/*
Generates a QR string for a user with the provided ID, as a URI
*/
func GetQrInfo(id string) (string, error) {
	_, err := GetUserInfo(id, nil)

	if err != nil {
		return "", errors.New("User does not exist.")
	}

	// Construct the URI

	uri, err := url.Parse("hackillinois://user")

	if err != nil {
		return "", err
	}

	// Generate JWT

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(time.Second * 20).Unix(),
		"userId": id,
	})

	signed_token, err := token.SignedString([]byte(common_config.TOKEN_SECRET))

	if err != nil {
		return "", err
	}

	// All the fields that will be embedded in the QR code URI
	parameters := url.Values{
		"userToken": []string{signed_token},
	}

	uri.RawQuery = parameters.Encode()

	return uri.String(), nil
}

/*
Returns all user stats
*/
func GetStats() (map[string]interface{}, error) {
	return db.GetStats("info", []string{}, nil)
}
