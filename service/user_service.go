package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/HackIllinois/api-auth/config"
	"github.com/HackIllinois/api-auth/models"
)

/*
	Send basic user info to the user service
*/
func SendUserInfo(id string, username string, email string) error {
	user_info := models.UserInfo{
		ID:       id,
		Username: username,
		Email:    email,
	}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&user_info)

	resp, err := http.Post(config.USER_SERVICE+"/user/", "application/json", &body)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("User service failed to update")
	}

	return nil
}

/*
	Given a user ID, fetch the user info corresponding to the ID.
*/
func GetUserInfo(id string) (models.UserInfo, error) {
	apiUserUrl := fmt.Sprintf(config.USER_SERVICE+"/user/%s/", id)
	resp, err := http.Get(apiUserUrl)
	var userInfo models.UserInfo
	json.NewDecoder(resp.Body).Decode(&userInfo)
	return userInfo, err
}
