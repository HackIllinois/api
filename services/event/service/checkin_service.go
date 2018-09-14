package service

import (
	"github.com/pattyjogal/api/services/event/config"
	"net/http"
)

/*
	Checks if the user has been checked in with the checkin service
*/
func IsUserCheckedIn(id string) (bool, error) {
	resp, err := http.Get(config.CHECKIN_SERVICE + "/checkin/" + id + "/")

	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}
