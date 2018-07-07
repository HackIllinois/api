package config

import (
	"github.com/arbor-dev/arbor/proxy"
	"github.com/arbor-dev/arbor/security"
	"os"
	"strconv"
)

var GATEWAY_PORT uint16

var TOKEN_SECRET = os.Getenv("TOKEN_SECRET")

var AUTH_SERVICE = os.Getenv("AUTH_SERVICE")
var USER_SERVICE = os.Getenv("USER_SERVICE")
var REGISTRATION_SERVICE = os.Getenv("REGISTRATION_SERVICE")
var DECISION_SERVICE = os.Getenv("DECISION_SERVICE")
var RSVP_SERVICE = os.Getenv("RSVP_SERVICE")
var CHECKIN_SERVICE = os.Getenv("CHECKIN_SERVICE")
var UPLOAD_SERVICE = os.Getenv("UPLOAD_SERVICE")
var MAIL_SERVICE = os.Getenv("MAIL_SERVICE")

func init() {
	port, err := strconv.ParseUint(os.Getenv("GATEWAY_PORT"), 10, 16)

	if err != nil {
		panic(err)
	}

	GATEWAY_PORT = uint16(port)
}

func LoadArborConfig() {
	security.AccessLogLocation = "log/access.log"
	security.ClientRegistryLocation = "clients.db"
	proxy.AccessControlPolicy = "*"
}
