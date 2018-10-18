package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
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

/*
	Add user with given id, to the specified mailing list.
	If the mailing list does not exist, it creates a new list with the user.
*/
func AddUserToMailList(user_id string, mail_list_id string) error {
	add_to_mail_list_url := fmt.Sprintf("%s/mail/list/add/", config.MAIL_SERVICE)

	mail_list := models.MailList{
		ID:      mail_list_id,
		UserIDs: []string{user_id},
	}

	update_request_body := bytes.Buffer{}
	json.NewEncoder(&update_request_body).Encode(&mail_list)

	content_type := "application/json"

	_, err := http.Post(add_to_mail_list_url, content_type, &update_request_body)

	if err != nil {
		// The mailing list didn't exist
		create_mail_list_url := fmt.Sprintf("%s/mail/list/create/", config.MAIL_SERVICE)

		create_request_body := bytes.Buffer{}
		json.NewEncoder(&create_request_body).Encode(&mail_list)

		_, err := http.Post(create_mail_list_url, content_type, &create_request_body)

		return err
	}

	return err
}
