package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"github.com/HackIllinois/api/common/datastore"
	"os"
)

var RSVP_DB_HOST string
var RSVP_DB_NAME string

var RSVP_PORT string

var AUTH_SERVICE string
var REGISTRATION_SERVICE string
var DECISION_SERVICE string
var MAIL_SERVICE string

var RSVP_DEFINITION datastore.DataStoreDefinition

var RSVP_STAT_FIELDS []string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	RSVP_DB_HOST, err = cfg_loader.Get("RSVP_DB_HOST")

	if err != nil {
		return err
	}

	RSVP_DB_NAME, err = cfg_loader.Get("RSVP_DB_NAME")

	if err != nil {
		return err
	}

	RSVP_PORT, err = cfg_loader.Get("RSVP_PORT")

	if err != nil {
		return err
	}

	AUTH_SERVICE, err = cfg_loader.Get("AUTH_SERVICE")

	if err != nil {
		return err
	}

	REGISTRATION_SERVICE, err = cfg_loader.Get("REGISTRATION_SERVICE")

	if err != nil {
		return err
	}

	DECISION_SERVICE, err = cfg_loader.Get("DECISION_SERVICE")

	if err != nil {
		return err
	}

	MAIL_SERVICE, err = cfg_loader.Get("MAIL_SERVICE")

	if err != nil {
		return err
	}

	err = cfg_loader.ParseInto("RSVP_DEFINITION", &RSVP_DEFINITION)

	if err != nil {
		return err
	}

	err = cfg_loader.ParseInto("RSVP_STAT_FIELDS", &RSVP_STAT_FIELDS)

	if err != nil {
		return err
	}

	return nil
}
