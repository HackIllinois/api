package service

import (
	"net/http"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/checkin/config"
)

/*
	Returns true if the user with specified id is registered, and false if not.
*/
func IsUserRegistered(id string) (bool, error) {
	status, err := apirequest.Get(config.REGISTRATION_SERVICE+"/registration/attendee/"+id+"/", nil)

	if err != nil {
		return false, err
	}

	return status == http.StatusOK, nil
}
