package service

import (
	"bytes"
	"encoding/json"
	"errors"
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

	client := http.Client{}
	req, err := http.NewRequest("POST", config.DECISION_SERVICE+"/decision/", &body)

	if err != nil {
		return err
	}

	req.Header.Set("HackIllinois-Identity", "registrationservice")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Decision service failed to create decision")
	}

	return nil
}
