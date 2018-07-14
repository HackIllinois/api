package service

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	
	"github.com/HackIllinois/api-registration/config"
	"github.com/HackIllinois/api-registration/models"
)

/*
	Send user with specified id a confirmation email, with template as specified. 
*/
func SendUserMail(id string, template string) error {
	api_mail_url := fmt.Sprintf("%s/send/", config.MAIL_SERVICE)

	mail_order := &models.MailOrder{
		IDs: []string{id},
		Template: template,
	}

	request_body, err := json.Marshal(mail_order)

	if err != nil {
		return err
	}

	content_type := "application/json"

	_, err = http.Post(api_mail_url, content_type, bytes.NewBuffer(request_body))

	if err != nil {
		return err
	}

	return nil
}