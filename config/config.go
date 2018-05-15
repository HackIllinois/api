package config

import (
	"os"
)

var GITHUB_CLIENT_ID = os.Getenv("GITHUB_CLIENT_ID")
var GITHUB_CLIENT_SECRET = os.Getenv("GITHUB_CLIENT_SECRET")
var TOKEN_SECRET = []byte(string(os.Getenv("TOKEN_SECRET")))

var AUTH_DB_HOST = os.Getenv("AUTH_DB_HOST")
var AUTH_DB_NAME = os.Getenv("AUTH_DB_NAME")

var AUTH_PORT = os.Getenv("AUTH_PORT")

var USER_SERVICE = os.Getenv("USER_SERVICE")
