package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var TOKEN_SECRET []byte

var AUTH_REDIRECT_URI string

var GITHUB_CLIENT_ID string
var GITHUB_CLIENT_SECRET string

var GOOGLE_CLIENT_ID string
var GOOGLE_CLIENT_SECRET string

var LINKEDIN_CLIENT_ID string
var LINKEDIN_CLIENT_SECRET string

var AUTH_DB_HOST string
var AUTH_DB_NAME string

var AUTH_PORT string

var USER_SERVICE string

var STAFF_DOMAIN string
var SYSTEM_ADMIN_EMAIL string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	secret, err := cfg_loader.Get("TOKEN_SECRET")

	if err != nil {
		return err
	}

	TOKEN_SECRET = []byte(secret)

	AUTH_REDIRECT_URI, err = cfg_loader.Get("AUTH_REDIRECT_URI")

	if err != nil {
		return err
	}

	GITHUB_CLIENT_ID, err = cfg_loader.Get("GITHUB_CLIENT_ID")

	if err != nil {
		return err
	}

	GITHUB_CLIENT_SECRET, err = cfg_loader.Get("GITHUB_CLIENT_SECRET")

	if err != nil {
		return err
	}

	GOOGLE_CLIENT_ID, err = cfg_loader.Get("GOOGLE_CLIENT_ID")

	if err != nil {
		return err
	}

	GOOGLE_CLIENT_SECRET, err = cfg_loader.Get("GOOGLE_CLIENT_SECRET")

	if err != nil {
		return err
	}

	LINKEDIN_CLIENT_ID, err = cfg_loader.Get("LINKEDIN_CLIENT_ID")

	if err != nil {
		return err
	}

	LINKEDIN_CLIENT_SECRET, err = cfg_loader.Get("LINKEDIN_CLIENT_SECRET")

	if err != nil {
		return err
	}

	AUTH_DB_HOST, err = cfg_loader.Get("AUTH_DB_HOST")

	if err != nil {
		return err
	}

	AUTH_DB_NAME, err = cfg_loader.Get("AUTH_DB_NAME")

	if err != nil {
		return err
	}

	AUTH_PORT, err = cfg_loader.Get("AUTH_PORT")

	if err != nil {
		return err
	}

	USER_SERVICE, err = cfg_loader.Get("USER_SERVICE")

	if err != nil {
		return err
	}

	STAFF_DOMAIN, err = cfg_loader.Get("STAFF_DOMAIN")

	if err != nil {
		return err
	}

	SYSTEM_ADMIN_EMAIL, err = cfg_loader.Get("SYSTEM_ADMIN_EMAIL")

	if err != nil {
		return err
	}

	return nil
}
