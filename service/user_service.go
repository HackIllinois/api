package service

import (
	"github.com/HackIllinois/api-user/database"
	"github.com/HackIllinois/api-user/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	Returns the info associated with the given user id
*/
func GetUserInfo(id string) (*models.UserInfo, error) {
	query := bson.M{
		"id": id,
	}

	var user_info models.UserInfo
	err := database.FindOne("info", query, &user_info)

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

	err := database.Update("info", selector, &user_info)

	if err == mgo.ErrNotFound {
		err = database.Insert("info", &user_info)
	}

	return err
}
