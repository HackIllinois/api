package config

import (
	"os"
)

var CHECKIN_DB_HOST = os.Getenv("CHECKIN_DB_HOST")
var CHECKIN_DB_NAME = os.Getenv("CHECKIN_DB_NAME")

var CHECKIN_PORT = os.Getenv("CHECKIN_PORT")
