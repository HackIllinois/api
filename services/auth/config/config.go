package config

import (
	"os"
	"github.com/HackIllinois/api/common/configloader"
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

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	secret, err := cfg_loader.Get("TOKEN_SECRET")

	if err != nil {
		panic(err)
	}

	TOKEN_SECRET = []byte(secret)

	AUTH_REDIRECT_URI, err = cfg_loader.Get("AUTH_REDIRECT_URI")

	if err != nil {
		panic(err)
	}

	GITHUB_CLIENT_ID, err = cfg_loader.Get("GITHUB_CLIENT_ID")

	if err != nil {
		panic(err)
	}

	GITHUB_CLIENT_SECRET, err = cfg_loader.Get("GITHUB_CLIENT_SECRET")

	if err != nil {
		panic(err)
	}

	GOOGLE_CLIENT_ID, err = cfg_loader.Get("GOOGLE_CLIENT_ID")

	if err != nil {
		panic(err)
	}

	GOOGLE_CLIENT_SECRET, err = cfg_loader.Get("GOOGLE_CLIENT_SECRET")

	if err != nil {
		panic(err)
	}

	LINKEDIN_CLIENT_ID, err = cfg_loader.Get("LINKEDIN_CLIENT_ID")

	if err != nil {
		panic(err)
	}

	LINKEDIN_CLIENT_SECRET, err = cfg_loader.Get("LINKEDIN_CLIENT_SECRET")

	if err != nil {
		panic(err)
	}

	AUTH_DB_HOST, err = cfg_loader.Get("AUTH_DB_HOST")

	if err != nil {
		panic(err)
	}

	AUTH_DB_NAME, err = cfg_loader.Get("AUTH_DB_NAME")

	if err != nil {
		panic(err)
	}

	AUTH_PORT, err = cfg_loader.Get("AUTH_PORT")

	if err != nil {
		panic(err)
	}

	USER_SERVICE, err = cfg_loader.Get("USER_SERVICE")

	if err != nil {
		panic(err)
	}
}
