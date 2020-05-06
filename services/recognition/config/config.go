package config

import (
	"os"

	"github.com/HackIllinois/api/common/configloader"
)

var RECOGNITION_DB_HOST string
var RECOGNITION_DB_NAME string

var RECOGNITION_PORT string


func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	RECOGNITION_DB_HOST, err = cfg_loader.Get("RECOGNITION_DB_HOST")

	if err != nil {
		return err
	}

	RECOGNITION_DB_NAME, err = cfg_loader.Get("RECOGNITION_DB_NAME")

	if err != nil {
		return err
	}

	RECOGNITION_PORT, err = cfg_loader.Get("RECOGNITION_PORT")

	if err != nil {
		return err
	}

	return nil
}
