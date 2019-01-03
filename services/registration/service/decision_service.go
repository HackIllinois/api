package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
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

	req, err := http.NewRequest("POST", config.DECISION_SERVICE+"/decision/", &body)

	if err != nil {
		return err
	}

	req.Header.Set("HackIllinois-Identity", "registrationservice")

	status, err := apirequest.Do(req, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Decision service failed to create decision.")
	}

	return nil
}
