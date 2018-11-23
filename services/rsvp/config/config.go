package config

import (
	"os"
	"github.com/HackIllinois/api/common/configloader"
)

var RSVP_DB_HOST string
var RSVP_DB_NAME string

var RSVP_PORT string

var AUTH_SERVICE string
var DECISION_SERVICE string
var MAIL_SERVICE string

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	RSVP_DB_HOST, err = cfg_loader.Get("RSVP_DB_HOST")

	if err != nil {
		panic(err)
	}

	RSVP_DB_NAME, err = cfg_loader.Get("RSVP_DB_NAME")

	if err != nil {
		panic(err)
	}

	RSVP_PORT, err = cfg_loader.Get("RSVP_PORT")

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
