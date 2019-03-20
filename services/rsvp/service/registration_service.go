package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/rsvp/config"
	"net/http"
)

/*
	Retrieve registration data from registration service
*/
func GetRegistrationData(id string) (map[string]interface{}, error) {
	registration_data := make(map[string]interface{})
	status, err := apirequest.Get(config.REGISTRATION_SERVICE+"/registration/"+id+"/", &registration_data)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Unable to retrieve data from registration service.")
	}

	return registration_data, nil
}
