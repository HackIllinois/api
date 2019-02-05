package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/decision/config"
	"github.com/HackIllinois/api/services/decision/models"
)

/*
	Send user with specified id a confirmation email, with template as specified.
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
		return errors.New("Error sending decision email.")
	}

	return nil
}

/*
	Adds user with specified id to an appropriate mail list, based on their current decision.
	If the mail list doesn't exist, a new one is created, containing the user.
*/
func AddUserToMailList(id string, decision *models.DecisionHistory) error {

	var mail_list_name string
	switch decision.Status {
	case "ACCEPTED":
		mail_list_name = fmt.Sprintf("accepted_wave_%v", decision.Wave)
	case "REJECTED":
		mail_list_name = "rejected"
	case "WAITLISTED":
		mail_list_name = "waitlisted"
	default:
		return errors.New("Decision status is not valid.")
	}

	mail_list := models.MailList{
		ID:      mail_list_name,
		UserIDs: []string{id},
	}

	status, err_update := apirequest.Post(config.MAIL_SERVICE+"/mail/list/add/", &mail_list, nil)

	if err_update == nil && status != http.StatusOK {
		// The mail list with given id does not exist.
		// A new one will be created with the current user in it.

		status, err_create := apirequest.Post(config.MAIL_SERVICE+"/mail/list/create/", &mail_list, nil)

		if err_create == nil && status != http.StatusOK {

			return errors.New(fmt.Sprintf("Failed to create new MailList with id %s.", mail_list_name))

		} else if err_create != nil {

			// Error creating / executing the create POST request.
			return err_create
		}
	}
	// If there was an error creating / executing the update POST request, it is returned.
	// Otherwise, the user should be in the correct mail list.
	return err_update
}
