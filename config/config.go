package config

import (
	"os"
)

var DECISION_DB_HOST = os.Getenv("DECISION_DB_HOST")
var DECISION_DB_NAME = os.Getenv("DECISION_DB_NAME")

var DECISION_PORT = os.Getenv("DECISION_PORT")

var MAIL_SERVICE = os.Getenv("MAIL_SERVICE")
