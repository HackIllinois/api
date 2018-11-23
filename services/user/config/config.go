package config

import (
	"os"
	"github.com/HackIllinois/api/common/configloader"
)

var USER_DB_HOST string
var USER_DB_NAME string

var USER_PORT string

func init() {

	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	USER_DB_HOST, err = cfg_loader.Get("USER_DB_HOST")

	if err != nil {
		panic(err)
	}

	USER_DB_NAME, err = cfg_loader.Get("USER_DB_NAME")

	if err != nil {
		panic(err)
	}

	USER_PORT, err = cfg_loader.Get("USER_PORT")

	if err != nil {
		panic(err)
	}
}
