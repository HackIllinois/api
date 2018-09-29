package service

import (
	"encoding/json"
	"errors"
	"github.com/ethan-lord/api/services/checkin/config"
	"github.com/ethan-lord/api/services/checkin/models"
	"net/http"
)

/*
	Checks if the user has been rsvped in the decision service
*/
func IsAttendeeRsvped(id string) (bool, error) {
	resp, err := http.Get(config.RSVP_SERVICE + "/rsvp/" + id + "/")

	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, errors.New("Rsvp service failed to return status")
	}

	var rsvp models.UserRsvp
	json.NewDecoder(resp.Body).Decode(&rsvp)

	return rsvp.IsAttending, nil
}
