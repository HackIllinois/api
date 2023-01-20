package config

import (
	"os"

	"github.com/HackIllinois/api/common/configloader"
)

var USER_DB_HOST string
var USER_DB_NAME string

var USER_PORT string

func Initialize() error {

	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	USER_DB_HOST, err = cfg_loader.Get("USER_DB_HOST")

	if err != nil {
		return err
	}

	USER_DB_NAME, err = cfg_loader.Get("USER_DB_NAME")

	if err != nil {
		return err
	}

	USER_PORT, err = cfg_loader.Get("USER_PORT")

	if err != nil {
		return err
	}

	return nil
}
