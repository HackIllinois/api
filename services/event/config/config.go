package config

import (
	"os"
	"github.com/HackIllinois/api/common/configloader"
)

var EVENT_DB_HOST string
var EVENT_DB_NAME string

var EVENT_PORT string

var CHECKIN_SERVICE string

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	EVENT_DB_HOST, err = cfg_loader.Get("EVENT_DB_HOST")

	if err != nil {
		panic(err)
	}

	EVENT_DB_NAME, err = cfg_loader.Get("EVENT_DB_NAME")

	if err != nil {
		panic(err)
	}

	EVENT_PORT, err = cfg_loader.Get("EVENT_PORT")

	if err != nil {
		panic(err)
	}

	CHECKIN_SERVICE, err = cfg_loader.Get("CHECKIN_SERVICE")

	if err != nil {
		panic(err)
	}
}
