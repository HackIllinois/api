package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/mail/config"
	"github.com/HackIllinois/api/services/mail/models"
	"net/http"
)

/*
	Get basic registration info belonging to id
*/
func GetRegistrationInfo(id string) (*models.AllRegistration, error) {
	var registration_info models.AllRegistration
	status, err := apirequest.Get(config.REGISTRATION_SERVICE+"/registration/"+id+"/", &registration_info)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Registration service failed to return information")
	}

	return &registration_info, nil
}
