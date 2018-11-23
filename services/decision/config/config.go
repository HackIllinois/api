package config

import (
	"os"
	"github.com/HackIllinois/api/common/configloader"
)

var DECISION_DB_HOST string
var DECISION_DB_NAME string

var DECISION_PORT string

var MAIL_SERVICE string

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	DECISION_DB_HOST, err = cfg_loader.Get("DECISION_DB_HOST")

	if err != nil {
		panic(err)
	}

	DECISION_DB_NAME, err = cfg_loader.Get("DECISION_DB_NAME")

	if err != nil {
		panic(err)
	}

	DECISION_PORT, err = cfg_loader.Get("DECISION_PORT")

	if err != nil {
		panic(err)
	}

	MAIL_SERVICE, err = cfg_loader.Get("MAIL_SERVICE")

	if err != nil {
		panic(err)
	}
}
