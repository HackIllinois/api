package config

import (
	"os"
)

var IS_PRODUCTION = os.Getenv("IS_PRODUCTION") == true
var PRINT_PORT = os.Getenv("PRINT_PORT")
var SNS_REGION = os.Getenv("AWS_REGION")
var PRINT_TOPIC = os.Getenv("PRINT_TOPIC")
