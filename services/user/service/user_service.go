package service

import (
	"errors"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.USER_DB_HOST, config.USER_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns the info associated with the given user id
*/
func GetUserInfo(id string) (*models.UserInfo, error) {
	query := bson.M{
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
	selector := bson.M{
		"id": id,
	}

	err := db.Update("info", selector, &user_info)

	if err == mgo.ErrNotFound {
		err = db.Insert("info", &user_info)
	}

	return err
}

/*
	Returns the users associated with the given parameters
*/
func GetFilteredUserInfo(parameters map[string][]string) (*models.FilteredUsers, error) {
	query := make(map[string]interface{})

	for key, values := range parameters {
		if len(values) > 1 {
			return nil, errors.New("Multiple usage of key " + key)
		}

		key = strings.ToLower(key)
		query[key] = bson.M{"$in": strings.Split(values[0], ",")}
	}

	var filtered_users models.FilteredUsers
	err := db.FindAll("info", query, &filtered_users.Users)
	if err != nil {
		return nil, err
	}

	return &filtered_users, nil
}
