package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/models"
	"net/http"
	"time"
)

/*
	Checks if the user has been accepted in the decision service
*/
func IsApplicantAcceptedAndActive(id string) (bool, bool, error) {
	var decision models.UserDecision
	status, err := apirequest.Get(config.DECISION_SERVICE+"/decision/"+id+"/", &decision)

	if err != nil {
		return false, false, err
	}

	if status != http.StatusOK {
		return false, false, errors.New("Decision service failed to return status.")
	}

	return decision.Status == "ACCEPTED" && decision.Finalized, time.Now().Unix() < decision.ExpiresAt, nil
}
