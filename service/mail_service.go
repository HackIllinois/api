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
	Send mail the the users with the given ids, using the provided template
	Substitution will be generated based on user info
*/
func SendMailByID(mail_order models.MailOrder) (*models.MailStatus, error) {
	var mail_info models.MailInfo

	mail_info.Content = models.Content{
		TemplateID: mail_order.Template,
	}

	mail_info.Recipients = make([]models.Recipient, len(mail_order.IDs))
	for i, id := range mail_order.IDs {
		user_info, err := GetUserInfo(id)

		if err != nil {
			return nil, err
		}

		mail_info.Recipients[i].Address = models.Address{
			Email: user_info.Email,
			Name:  user_info.Username,
		}
		mail_info.Recipients[i].Substitutions = models.Substitutions{
			"name": user_info.Username,
		}
	}

	return SendMail(mail_info)
}

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
