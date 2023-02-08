package config

import (
	"os"

	"github.com/HackIllinois/api/common/configloader"
)

var (
	EVENT_DB_HOST string
	EVENT_DB_NAME string
)

var EVENT_PORT string

var CHECKIN_SERVICE string

var PROFILE_SERVICE string

var RSVP_SERVICE string

var EVENT_CHECKIN_TIME_RESTRICTED bool

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

	PROFILE_SERVICE, err = cfg_loader.Get("PROFILE_SERVICE")

	if err != nil {
		return err
	}

	RSVP_SERVICE, err = cfg_loader.Get("RSVP_SERVICE")

	if err != nil {
		return err
	}

	checkin_time_res_str, err := cfg_loader.Get("EVENT_CHECKIN_TIME_RESTRICTED")
	if err != nil {
		return err
	}

	EVENT_CHECKIN_TIME_RESTRICTED = (checkin_time_res_str == "true")

	return nil
}
