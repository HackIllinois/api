package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var IS_PRODUCTION bool

var MAIL_DB_HOST string
var MAIL_DB_NAME string

var MAIL_PORT string

var SPARKPOST_API string
var SPARKPOST_APIKEY string

var USER_SERVICE string
var REGISTRATION_SERVICE string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	MAIL_DB_HOST, err = cfg_loader.Get("MAIL_DB_HOST")

	if err != nil {
		return err
	}

	MAIL_DB_NAME, err = cfg_loader.Get("MAIL_DB_NAME")

	if err != nil {
		return err
	}

	MAIL_PORT, err = cfg_loader.Get("MAIL_PORT")

	if err != nil {
		return err
	}

	SPARKPOST_API, err = cfg_loader.Get("SPARKPOST_API")

	if err != nil {
		return err
	}

	SPARKPOST_APIKEY, err = cfg_loader.Get("SPARKPOST_APIKEY")

	if err != nil {
		return err
	}

	USER_SERVICE, err = cfg_loader.Get("USER_SERVICE")

	if err != nil {
		return err
	}

	REGISTRATION_SERVICE, err = cfg_loader.Get("REGISTRATION_SERVICE")

	if err != nil {
		return err
	}

	production, err := cfg_loader.Get("IS_PRODUCTION")

	if err != nil {
		return err
	}

	IS_PRODUCTION = (production == "true")

	return nil
}
