package service

import (
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api/services/print/config"
	"github.com/HackIllinois/api/services/print/models"
	"net/http"
)

/*
	Get basic user info from user serivce
*/
var GetUserInfo = func(id string) (*models.UserInfo, error) {
	resp, err := http.Get(config.USER_SERVICE + "/user/" + id + "/")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("User service failed to return information")
	}

	var user_info models.UserInfo
	json.NewDecoder(resp.Body).Decode(&user_info)

	return &user_info, nil
}
