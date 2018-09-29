package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethan-lord/api/services/auth/config"
	"github.com/ethan-lord/api/services/auth/models"
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
func GetUserInfo(id string) (*models.UserInfo, error) {
	api_user_url := fmt.Sprintf("%s/user/%s/", config.USER_SERVICE, id)
	client := &http.Client{}
	req, err := http.NewRequest("GET", api_user_url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("GET request to api-user(%s) resulted in a response with a non-200 status code.", api_user_url))
	}

	defer resp.Body.Close()

	var user_info models.UserInfo
	json.NewDecoder(resp.Body).Decode(&user_info)
	return &user_info, err
}
