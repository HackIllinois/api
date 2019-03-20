package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/models"
	"net/http"
)

/*
	Checks if the user has been rsvped in the decision service
*/
func IsAttendeeRsvped(id string) (bool, error) {
	var rsvp models.UserRsvp
	status, err := apirequest.Get(config.RSVP_SERVICE+"/rsvp/"+id+"/", &rsvp)

	if err != nil {
		return false, err
	}

	if status != http.StatusOK {
		return false, errors.New("Rsvp service failed to return status")
	}

	return rsvp.IsAttending, nil
}

/*
	Retrieve rsvp data from rsvp service
*/
func GetRsvpData(id string) (map[string]interface{}, error) {
	rsvp_data := make(map[string]interface{})
	status, err := apirequest.Get(config.RSVP_SERVICE+"/rsvp/"+id+"/", &rsvp_data)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Unable to retrieve data from rsvp service.")
	}

	return rsvp_data, nil
}
