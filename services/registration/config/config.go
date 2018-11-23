package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var REGISTRATION_DB_HOST string
var REGISTRATION_DB_NAME string

var REGISTRATION_PORT string

var USER_SERVICE string
var AUTH_SERVICE string
var DECISION_SERVICE string
var MAIL_SERVICE string

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	REGISTRATION_DB_HOST, err = cfg_loader.Get("REGISTRATION_DB_HOST")

	if err != nil {
		panic(err)
	}

	REGISTRATION_DB_NAME, err = cfg_loader.Get("REGISTRATION_DB_NAME")

	if err != nil {
		panic(err)
	}

	REGISTRATION_PORT, err = cfg_loader.Get("REGISTRATION_PORT")

	if err != nil {
		panic(err)
	}

	USER_SERVICE, err = cfg_loader.Get("USER_SERVICE")

	if err != nil {
		panic(err)
	}

	AUTH_SERVICE, err = cfg_loader.Get("AUTH_SERVICE")

	if err != nil {
		panic(err)
	}

	DECISION_SERVICE, err = cfg_loader.Get("DECISION_SERVICE")

	if err != nil {
		panic(err)
	}

	MAIL_SERVICE, err = cfg_loader.Get("MAIL_SERVICE")

	if err != nil {
		panic(err)
	}
}
