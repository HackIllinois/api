package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
)

/*
	Send user with specified id a confirmation email, with template as specified.
	If there is an error sending the confirmation email, the registration is failed.
*/
func SendUserMail(id string, template string) error {
	mail_order := models.MailOrder{
		IDs:      []string{id},
		Template: template,
	}

	status, err := apirequest.Post(config.MAIL_SERVICE+"/mail/send/", &mail_order, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New("Error sending confirmation email, therefore, registration was failed.")
	}

	return nil
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

	status, err := apirequest.Post(add_to_mail_list_url, &mail_list, nil)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		// The mailing list didn't exist
		create_mail_list_url := fmt.Sprintf("%s/mail/list/create/", config.MAIL_SERVICE)

		status, err = apirequest.Post(create_mail_list_url, &mail_list, nil)

		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.New("Mailing list does not exist, furthermore, creation of mailing list failed.")
		}

		return nil
	}

	return nil
}
