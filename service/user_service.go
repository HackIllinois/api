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
func GetUserInfo(id string, tokenString string) (models.UserInfo, error) {
	api_user_url := fmt.Sprintf("http://localhost:8000/user/%s/", id)
	client := &http.Client{}
	req, err := http.NewRequest("GET", api_user_url, nil)

	if err != nil {
		errors.New("GET request to api-user failed to be created")
	}

	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	var user_info models.UserInfo

	if err != nil {
		return user_info, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&user_info)
	return user_info, err
}
