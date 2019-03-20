package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
	"strconv"
)

var DECISION_DB_HOST string
var DECISION_DB_NAME string

var DECISION_PORT string

var MAIL_SERVICE string

var DECISION_EXPIRATION_HOURS int

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	DECISION_DB_HOST, err = cfg_loader.Get("DECISION_DB_HOST")

	if err != nil {
		return err
	}

	DECISION_DB_NAME, err = cfg_loader.Get("DECISION_DB_NAME")

	if err != nil {
		return err
	}

	DECISION_PORT, err = cfg_loader.Get("DECISION_PORT")

	if err != nil {
		return err
	}

	MAIL_SERVICE, err = cfg_loader.Get("MAIL_SERVICE")

	if err != nil {
		return err
	}

	decision_expiration_str, err := cfg_loader.Get("DECISION_EXPIRATION_HOURS")

	if err != nil {
		return err
	}

	decision_expiration, err := strconv.Atoi(decision_expiration_str)

	if err != nil {
		return err
	}

	DECISION_EXPIRATION_HOURS = decision_expiration

	return nil
}
