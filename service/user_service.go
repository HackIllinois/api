package service

import (
	"github.com/HackIllinois/api-commons/database"
	"github.com/HackIllinois/api-user/config"
	"github.com/HackIllinois/api-user/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
