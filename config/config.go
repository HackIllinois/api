package config

import (
	"os"
)

var EVENT_DB_HOST = os.Getenv("EVENT_DB_HOST")
var EVENT_DB_NAME = os.Getenv("EVENT_DB_NAME")

var EVENT_PORT = os.Getenv("EVENT_PORT")

var CHECKIN_SERVICE = os.Getenv("CHECKIN_SERVICE")
