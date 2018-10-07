package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"github.com/HackIllinois/api/common/apirequest"
	"net/http"
)

/*
	Create the initial decision for the application with the decision service
*/
func AddInitialDecision(id string) error {
	decision := models.UserDecision{
		ID:     id,
		Status: "PENDING",
	}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&decision)

	status, err := apirequest.Post(config.DECISION_SERVICE + "/decision/", &body, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Decision service failed to create decision")
	}

	return nil
}
