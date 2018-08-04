package config

import (
	"os"
)

var STAT_DB_HOST = os.Getenv("STAT_DB_HOST")
var STAT_DB_NAME = os.Getenv("STAT_DB_NAME")

var STAT_PORT = os.Getenv("STAT_PORT")
