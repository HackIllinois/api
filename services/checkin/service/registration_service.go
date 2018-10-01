package service

import (
	"fmt"
	"net/http"

	"github.com/HackIllinois/api/services/checkin/config"
)

/*
	Returns true if the user with specified id is registered, and false if not.
*/
func IsUserRegistered(id string) (bool, error) {
	api_registration_url := fmt.Sprintf("%s/registration/attendee/%s/", config.REGISTRATION_SERVICE, id)

	resp, err := http.Get(api_registration_url)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}
