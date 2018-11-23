package config

import (
	"os"
	"github.com/HackIllinois/api/common/configloader"
)

var IS_PRODUCTION bool

var MAIL_DB_HOST string
var MAIL_DB_NAME string

var MAIL_PORT string

var SPARKPOST_API string
var SPARKPOST_APIKEY string

var USER_SERVICE string

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	MAIL_DB_HOST, err = cfg_loader.Get("MAIL_DB_HOST")

	if err != nil {
		panic(err)
	}

	MAIL_DB_NAME, err = cfg_loader.Get("MAIL_DB_NAME")

	if err != nil {
		panic(err)
	}

	MAIL_PORT, err = cfg_loader.Get("MAIL_PORT")

	if err != nil {
		panic(err)
	}

	SPARKPOST_API, err = cfg_loader.Get("SPARKPOST_API")

	if err != nil {
		panic(err)
	}

	SPARKPOST_APIKEY, err = cfg_loader.Get("SPARKPOST_APIKEY")

	if err != nil {
		panic(err)
	}

	USER_SERVICE, err = cfg_loader.Get("USER_SERVICE")

	if err != nil {
		panic(err)
	}

	production, err := cfg_loader.Get("IS_PRODUCTION")

	if err != nil {
		panic(err)
	}

	IS_PRODUCTION = (production == "true")
}
