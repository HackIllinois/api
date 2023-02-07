package service

import (
	"errors"
	"net/http"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/event/config"
)

func GetRsvpData(id string) (map[string]interface{}, error) {
	var rsvp_data map[string]interface{}
	status, err := apirequest.Get(config.RSVP_SERVICE+"/rsvp/"+id+"/", &rsvp_data)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Getting RSVP data failed")
	}

	return rsvp_data, nil
}
