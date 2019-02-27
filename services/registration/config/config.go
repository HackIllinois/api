package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"github.com/HackIllinois/api/common/datastore"
	"os"
)

var REGISTRATION_DB_HOST string
var REGISTRATION_DB_NAME string

var REGISTRATION_PORT string

var USER_SERVICE string
var AUTH_SERVICE string
var DECISION_SERVICE string
var MAIL_SERVICE string

var REGISTRATION_DEFINITION datastore.DataStoreDefinition
var MENTOR_REGISTRATION_DEFINITION datastore.DataStoreDefinition

var REGISTRATION_STAT_FIELDS []string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	REGISTRATION_DB_HOST, err = cfg_loader.Get("REGISTRATION_DB_HOST")

	if err != nil {
		return err
	}

	REGISTRATION_DB_NAME, err = cfg_loader.Get("REGISTRATION_DB_NAME")

	if err != nil {
		return err
	}

	REGISTRATION_PORT, err = cfg_loader.Get("REGISTRATION_PORT")

	if err != nil {
		return err
	}

	USER_SERVICE, err = cfg_loader.Get("USER_SERVICE")

	if err != nil {
		return err
	}

	AUTH_SERVICE, err = cfg_loader.Get("AUTH_SERVICE")

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

	err = cfg_loader.ParseInto("REGISTRATION_DEFINITION", &REGISTRATION_DEFINITION)

	if err != nil {
		return err
	}

	err = cfg_loader.ParseInto("MENTOR_REGISTRATION_DEFINITION", &MENTOR_REGISTRATION_DEFINITION)

	if err != nil {
		return err
	}

	err = cfg_loader.ParseInto("REGISTRATION_STAT_FIELDS", &REGISTRATION_STAT_FIELDS)

	if err != nil {
		return err
	}

	return nil
}
