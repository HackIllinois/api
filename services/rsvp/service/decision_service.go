package service

import (
	"errors"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/models"
	"github.com/HackIllinois/api/common/apirequest"
	"net/http"
)

/*
	Checks if the user has been accepted in the decision service
*/
func IsApplicantAccepted(id string) (bool, error) {
	var decision models.UserDecision
	status, err := apirequest.Get(config.DECISION_SERVICE + "/decision/" + id + "/", &decision)

	if err != nil {
		return false, err
	}

	if status != http.StatusOK {
		return false, errors.New("Decision service failed to return status")
	}

	return decision.Status == "ACCEPTED", nil
}
