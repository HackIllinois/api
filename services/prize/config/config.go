package config

import (
	"os"

	"github.com/HackIllinois/api/common/configloader"
)

var PRIZE_DB_HOST string
var PRIZE_DB_NAME string

var PRIZE_PORT string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	PRIZE_DB_HOST, err = cfg_loader.Get("PRIZE_DB_HOST")

	if err != nil {
		return err
	}

	PRIZE_DB_NAME, err = cfg_loader.Get("PRIZE_DB_NAME")

	if err != nil {
		return err
	}

	PRIZE_PORT, err = cfg_loader.Get("PRIZE_PORT")

	if err != nil {
		return err
	}

	return nil
}
