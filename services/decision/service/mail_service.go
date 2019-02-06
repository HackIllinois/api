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
	Gets the mailing list to add and remove from, based on a decision.
*/
func GetMailListFromDecision(decision *models.DecisionHistory) (string, error) {
	switch decision.Status {
	case "ACCEPTED":
		return fmt.Sprintf("accepted_wave_%v", decision.Wave), nil
	case "REJECTED":
		return "rejected", nil
	case "WAITLISTED":
		return "waitlisted", nil
	default:
		return "", errors.New("Decision status is not valid.")
	}
}

/*
	Adds user with specified id to an appropriate mail list, based on their current decision.
	If the mail list doesn't exist, a new one is created, containing the user.
*/
func AddUserToMailList(id string, decision *models.DecisionHistory) error {

	mail_list_name, err := GetMailListFromDecision(decision)

	if err != nil {
		return err
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

/*
	Removes user from appropriate mail list, based on decision.
*/
func RemoveUserFromMailList(id string, decision *models.DecisionHistory) error {
	mail_list_name, err := GetMailListFromDecision(decision)

	if err != nil {
		return err
	}

	mail_list := models.MailList{
		ID:      mail_list_name,
		UserIDs: []string{id},
	}

	status, err_remove := apirequest.Post(config.MAIL_SERVICE+"/mail/list/remove/", &mail_list, nil)

	if err_remove == nil && status != http.StatusOK {
		return errors.New(fmt.Sprintf("Failed to remove user from mailing list %s.", mail_list_name))
	}

	return err_remove
}
