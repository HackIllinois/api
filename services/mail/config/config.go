package config

import (
	"os"
)

var IS_PRODUCTION = (os.Getenv("IS_PRODUCTION") == "true")

var MAIL_DB_HOST = os.Getenv("MAIL_DB_HOST")
var MAIL_DB_NAME = os.Getenv("MAIL_DB_NAME")

var MAIL_PORT = os.Getenv("MAIL_PORT")

var SPARKPOST_API = os.Getenv("SPARKPOST_API")
var SPARKPOST_APIKEY = os.Getenv("SPARKPOST_APIKEY")

var USER_SERVICE = os.Getenv("USER_SERVICE")
