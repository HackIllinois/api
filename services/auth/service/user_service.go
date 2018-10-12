package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
)

/*
	Send basic user info to the user service
*/
func SendUserInfo(id string, username string, first_name string, last_name string, email string) error {
	user_info := models.UserInfo{
		ID:        id,
		Username:  username,
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
	}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_info)

	status, err := apirequest.Post(config.USER_SERVICE+"/user/", &body, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("User service failed to update")
	}

	return nil
}

/*
	Given a user ID, fetch the user info corresponding to the ID.
*/
func GetUserInfo(id string) (*models.UserInfo, error) {
	var user_info models.UserInfo
	status, err := apirequest.Get(config.USER_SERVICE+"/user/"+id+"/", &user_info)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Failed to retrieve info from user service")
	}

	return &user_info, err
}
