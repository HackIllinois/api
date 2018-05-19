package config

import (
	"os"
)

var RSVP_DB_HOST = os.Getenv("RSVP_DB_HOST")
var RSVP_DB_NAME = os.Getenv("RSVP_DB_NAME")

var RSVP_PORT = os.Getenv("RSVP_PORT")

var DECISION_SERVICE = os.Getenv("DECISION_SERVICE")
