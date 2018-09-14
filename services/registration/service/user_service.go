package service

import (
	"encoding/json"
	"errors"
	"github.com/pattyjogal/api/services/registration/config"
	"github.com/pattyjogal/api/services/registration/models"
	"net/http"
)

/*
	Get basic user info from user serivce
*/
func GetUserInfo(id string) (*models.UserInfo, error) {
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
