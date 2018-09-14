package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pattyjogal/api/services/registration/config"
	"github.com/pattyjogal/api/services/registration/models"
)

/*
	Send user with specified id a confirmation email, with template as specified.
*/
func SendUserMail(id string, template string) error {
	api_mail_url := fmt.Sprintf("%s/mail/send/", config.MAIL_SERVICE)

	mail_order := models.MailOrder{
		IDs:      []string{id},
		Template: template,
	}

	request_body := bytes.Buffer{}
	json.NewEncoder(&request_body).Encode(&mail_order)

	content_type := "application/json"

	_, err := http.Post(api_mail_url, content_type, &request_body)

	return err
}
