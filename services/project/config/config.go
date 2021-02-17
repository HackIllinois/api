package config

import (
	"os"

	"github.com/HackIllinois/api/common/configloader"
)

var PROJECT_DB_HOST string
var PROJECT_DB_NAME string

var PROJECT_PORT string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	PROJECT_DB_HOST, err = cfg_loader.Get("PROJECT_DB_HOST")

	if err != nil {
		return err
	}

	PROJECT_DB_NAME, err = cfg_loader.Get("PROJECT_DB_NAME")

	if err != nil {
		return err
	}

	PROJECT_PORT, err = cfg_loader.Get("PROJECT_PORT")

	if err != nil {
		return err
	}

	return nil
}
