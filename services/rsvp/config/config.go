package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"github.com/HackIllinois/api/common/datastore"
	"os"
	"strconv"
)

var RSVP_DB_HOST string
var RSVP_DB_NAME string

var RSVP_PORT string

var AUTH_SERVICE string
var DECISION_SERVICE string
var MAIL_SERVICE string

var DECISION_EXPIRATION_HOURS int

var RSVP_DEFINITION datastore.DataStoreDefinition

var RSVP_STAT_FIELDS []string

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

	decision_expiration_str, err := cfg_loader.Get("DECISION_EXPIRATION_HOURS")

	if err != nil {
		panic(err)
	}

	decision_expiration, err := strconv.Atoi(decision_expiration_str)

	if err != nil {
		panic(err)
	}

	DECISION_EXPIRATION_HOURS = decision_expiration

	err = cfg_loader.ParseInto("RSVP_DEFINITION", &RSVP_DEFINITION)

	if err != nil {
		panic(err)
	}

	err = cfg_loader.ParseInto("RSVP_STAT_FIELDS", &RSVP_STAT_FIELDS)

	if err != nil {
		panic(err)
	}
}
