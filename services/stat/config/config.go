package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var STAT_DB_HOST string
var STAT_DB_NAME string

var STAT_PORT string

var STAT_ENDPOINTS map[string]string

func Initialize() error {

	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	STAT_DB_HOST, err = cfg_loader.Get("STAT_DB_HOST")

	if err != nil {
		return err
	}

	STAT_DB_NAME, err = cfg_loader.Get("STAT_DB_NAME")

	if err != nil {
		return err
	}

	STAT_PORT, err = cfg_loader.Get("STAT_PORT")

	if err != nil {
		return err
	}

	err = cfg_loader.ParseInto("STAT_ENDPOINTS", &STAT_ENDPOINTS)

	if err != nil {
		return err
	}

	return nil
}
