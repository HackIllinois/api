package config

import (
	"os"

	"github.com/HackIllinois/api/common/configloader"
)

var PROFILE_DB_HOST string
var PROFILE_DB_NAME string

var PROFILE_PORT string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	PROFILE_DB_HOST, err = cfg_loader.Get("PROFILE_DB_HOST")

	if err != nil {
		return err
	}

	PROFILE_DB_NAME, err = cfg_loader.Get("PROFILE_DB_NAME")

	if err != nil {
		return err
	}

	PROFILE_PORT, err = cfg_loader.Get("PROFILE_PORT")

	if err != nil {
		return err
	}

	return nil
}
