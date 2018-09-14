package service

import (
	"encoding/json"
	"errors"
	"github.com/pattyjogal/api/services/rsvp/config"
	"github.com/pattyjogal/api/services/rsvp/models"
	"net/http"
)

/*
	Checks if the user has been accepted in the decision service
*/
func IsApplicantAccepted(id string) (bool, error) {
	resp, err := http.Get(config.DECISION_SERVICE + "/decision/" + id + "/")

	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, errors.New("Decision service failed to return status")
	}

	var decision models.UserDecision
	json.NewDecoder(resp.Body).Decode(&decision)

	return decision.Status == "ACCEPTED", nil
}
