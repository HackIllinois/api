package config

import (
	"os"
)

var REGISTRATION_DB_HOST = os.Getenv("REGISTRATION_DB_HOST")
var REGISTRATION_DB_NAME = os.Getenv("REGISTRATION_DB_NAME")
