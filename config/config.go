package config

import (
	"os"
)

var GITHUB_CLIENT_ID = os.Getenv("GITHUB_CLIENT_ID")
var GITHUB_CLIENT_SECRET = os.Getenv("GITHUB_CLIENT_SECRET")
var TOKEN_SECRET = []byte(string(os.Getenv("TOKEN_SECRET")))
