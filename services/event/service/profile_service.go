package service

import (
	"net/http"
	"strconv"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/event/config"
)

/*
	Checks if the user has been checked in already
*/
func AlreadyRedeemedEvent(event_id string) (bool, error) {
	status, err := apirequest.Get(config.PROFILE_SERVICE+"/checkin/"+event_id+"/", nil)

	if err != nil {
		return false, err
	}

	return status == http.StatusOK, nil
}

/*
	Performs a get and a put operation on the profile to increment the current number of points
*/
func UpdatePoints(points int) (bool, error) {
	status, err := apirequest.Get(config.PROFILE_SERVICE+"/award/"+strconv.Itoa(points)+"/", nil)

	if err != nil {
		return false, err
	}

	return status == http.StatusOK, nil
}
