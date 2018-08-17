package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ReflectionsProjections/api/services/decision/config"
	"github.com/ReflectionsProjections/api/services/decision/models"
)

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

	request_body := bytes.Buffer{}
	json.NewEncoder(&request_body).Encode(&mail_list)

	// URL to update the MailList with new IDs
	api_mail_update_url := fmt.Sprintf("%s/mail/list/add/", config.MAIL_SERVICE)

	content_type := "application/json"

	resp, err_update := http.Post(api_mail_update_url, content_type, &request_body)

	if err_update == nil && resp.StatusCode != http.StatusOK {
		// The mail list with given id does not exist.
		// A new one will be created with the current user in it.
		api_mail_create_url := fmt.Sprintf("%s/mail/list/create/", config.MAIL_SERVICE)

		// Since the buffer gets consumed after the preceding POST request
		json.NewEncoder(&request_body).Encode(&mail_list)

		resp, err_create := http.Post(api_mail_create_url, content_type, &request_body)

		if err_create == nil && resp.StatusCode != http.StatusOK {

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
