package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var CHECKIN_DB_HOST string
var CHECKIN_DB_NAME string

var CHECKIN_PORT string

var RSVP_SERVICE string
var REGISTRATION_SERVICE string

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	CHECKIN_DB_HOST, err = cfg_loader.Get("CHECKIN_DB_HOST")

	if err != nil {
		panic(err)
	}

	CHECKIN_DB_NAME, err = cfg_loader.Get("CHECKIN_DB_NAME")

	if err != nil {
		panic(err)
	}

	CHECKIN_PORT, err = cfg_loader.Get("CHECKIN_PORT")

	if err != nil {
		panic(err)
	}

	RSVP_SERVICE, err = cfg_loader.Get("RSVP_SERVICE")

	if err != nil {
		panic(err)
	}

	REGISTRATION_SERVICE, err = cfg_loader.Get("REGISTRATION_SERVICE")

	if err != nil {
		panic(err)
	}
}
