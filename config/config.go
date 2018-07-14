package config

import (
	"os"
)

var TOKEN_SECRET = []byte(string(os.Getenv("TOKEN_SECRET")))

var AUTH_REDIRECT_URI = os.Getenv("AUTH_REDIRECT_URI")

var GITHUB_CLIENT_ID = os.Getenv("GITHUB_CLIENT_ID")
var GITHUB_CLIENT_SECRET = os.Getenv("GITHUB_CLIENT_SECRET")

var GOOGLE_CLIENT_ID = os.Getenv("GOOGLE_CLIENT_ID")
var GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")

var AUTH_DB_HOST = os.Getenv("AUTH_DB_HOST")
var AUTH_DB_NAME = os.Getenv("AUTH_DB_NAME")

var AUTH_PORT = os.Getenv("AUTH_PORT")

var USER_SERVICE = os.Getenv("USER_SERVICE")
