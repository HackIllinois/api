package service

import (
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/models"
)

/*
	Send user with specified id a confirmation email, with template as specified.
*/
func SendUserMail(id string, template string) error {
	mail_order := models.MailOrder{
		IDs:      []string{id},
		Template: template,
	}

	_, err := apirequest.Post(config.MAIL_SERVICE+"/mail/send/", &mail_order, nil)

	return err
}
