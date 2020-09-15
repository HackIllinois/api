package service

import (
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/event/config"
	"net/http"
)

/*
	Checks if the user has been checked in with the checkin service
*/
func IsUserCheckedIn(id string) (bool, error) {
	status, err := apirequest.Get(config.CHECKIN_SERVICE+"/checkin/"+id+"/", nil)

	if err != nil {
		return false, err
	}

	return status == http.StatusOK, nil
}
