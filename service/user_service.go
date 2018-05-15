package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api-auth/config"
	"github.com/HackIllinois/api-auth/models"
	"net/http"
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
