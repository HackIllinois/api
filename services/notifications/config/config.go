package config

import (
	"os"
)

var NOTIFICATIONS_DB_HOST = os.Getenv("NOTIFICATIONS_DB_HOST")
var NOTIFICATIONS_DB_NAME = os.Getenv("NOTIFICATIONS_DB_NAME")

var NOTIFICATIONS_PORT = os.Getenv("NOTIFICATIONS_PORT")

var IS_PRODUCTION = (os.Getenv("IS_PRODUCTION") == "true")

var SNS_REGION = os.Getenv("SNS_REGION")
