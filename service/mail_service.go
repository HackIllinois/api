package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api-mail/config"
	"github.com/HackIllinois/api-mail/models"
	"net/http"
)

/*
	Send mail based on the given mailing info
	Returns the results of sending the mail
*/
func SendMail(mail_info models.MailInfo) (*models.MailStatus, error) {
	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&mail_info)

	client := http.Client{}
	req, err := http.NewRequest("POST", config.SPARKPOST_API+"/transmissions/", &body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", config.SPARKPOST_APIKEY)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to send mail")
	}

	var mail_status models.MailStatus
	json.NewDecoder(resp.Body).Decode(&mail_status)

	return &mail_status, nil
}
