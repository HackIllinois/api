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

	status, err := apirequest.Post(add_to_mail_list_url, &update_request_body, nil)

	if err != nil || status != http.StatusOK {
		// The mailing list didn't exist
		create_mail_list_url := fmt.Sprintf("%s/mail/list/create/", config.MAIL_SERVICE)

		create_request_body := bytes.Buffer{}
		json.NewEncoder(&create_request_body).Encode(&mail_list)

		_, err := apirequest.Post(create_mail_list_url, &create_request_body, nil)

		return err
	}

	return err
}
