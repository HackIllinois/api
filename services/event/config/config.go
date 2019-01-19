package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var EVENT_DB_HOST string
var EVENT_DB_NAME string

var EVENT_PORT string

var CHECKIN_SERVICE string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	EVENT_DB_HOST, err = cfg_loader.Get("EVENT_DB_HOST")

	if err != nil {
		return err
	}

	EVENT_DB_NAME, err = cfg_loader.Get("EVENT_DB_NAME")

	if err != nil {
		return err
	}

	EVENT_PORT, err = cfg_loader.Get("EVENT_PORT")

	if err != nil {
		return err
	}

	CHECKIN_SERVICE, err = cfg_loader.Get("CHECKIN_SERVICE")

	if err != nil {
		return err
	}

	return nil
}
