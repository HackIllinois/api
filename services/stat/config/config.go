package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var STAT_DB_HOST string
var STAT_DB_NAME string

var STAT_PORT string

func init() {

	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	STAT_DB_HOST, err = cfg_loader.Get("STAT_DB_HOST")

	if err != nil {
		panic(err)
	}

	STAT_DB_NAME, err = cfg_loader.Get("STAT_DB_NAME")

	if err != nil {
		panic(err)
	}

	STAT_PORT, err = cfg_loader.Get("STAT_PORT")

	if err != nil {
		panic(err)
	}
}
