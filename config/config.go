package config

import (
	"os"
)

var USER_DB_HOST = os.Getenv("USER_DB_HOST")
var USER_DB_NAME = os.Getenv("USER_DB_NAME")
