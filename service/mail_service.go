package service

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"errors"

	"github.com/HackIllinois/api-decision/config"
	"github.com/HackIllinois/api-decision/models"
	"gopkg.in/mgo.v2"
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

	mail_list := &models.MailList{
		ID: mail_list_name,
		UserIDs: []string{id},
	}
	
	request_body, err := json.Marshal(mail_list)

	if err != nil {
		return err
	}

	request_body_buffer := bytes.NewBuffer(request_body)
	
	// URL to update the MailList with new IDs
	api_mail_update_url := fmt.Sprintf("%s/list/add/", config.MAIL_SERVICE)
	content_type := "application/json"
	_, err_update := http.Post(api_mail_update_url, content_type, request_body_buffer)

	if err_update == mgo.ErrNotFound {
		// The mail list with given id does not exist. 
		// A new one will be created with the current user in it.
		api_mail_create_url := fmt.Sprintf("%s/list/create/", config.MAIL_SERVICE)
		_, err_create := http.Post(api_mail_create_url, content_type, request_body_buffer)
		
		if err_create != nil {
			return err_create
		}

		return nil
	} else if err_update != nil {
		return err_update
	}
	return nil
}