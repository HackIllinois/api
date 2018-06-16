package config

import (
	"os"
)

var MAIL_DB_HOST = os.Getenv("MAIL_DB_HOST")
var MAIL_DB_NAME = os.Getenv("MAIL_DB_NAME")

var MAIL_PORT = os.Getenv("MAIL_PORT")

var SPARKPOST_API = os.Getenv("SPARKPOST_API")
var SPARKPOST_APIKEY = os.Getenv("SPARKPOST_APIKEY")
