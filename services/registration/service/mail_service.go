package service

import (
	"bytes"
	"encoding/json"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
)

/*
	Send user with specified id a confirmation email, with template as specified.
*/
func SendUserMail(id string, template string) error {
	mail_order := models.MailOrder{
		IDs:      []string{id},
		Template: template,
	}

	request_body := bytes.Buffer{}
	json.NewEncoder(&request_body).Encode(&mail_order)

	_, err := apirequest.Post(config.MAIL_SERVICE+"/mail/send/", &request_body, nil)

	return err
}
