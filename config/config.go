package config

import (
	"github.com/arbor-dev/arbor/proxy"
	"github.com/arbor-dev/arbor/security"
	"os"
)

var TEST_SERVICE = os.Getenv("TEST_SERVICE")
var AUTH_SERVICE = os.Getenv("AUTH_SERVICE")
var USER_SERVICE = os.Getenv("USER_SERVICE")
var REGISTRATION_SERVICE = os.Getenv("REGISTRATION_SERVICE")
var DECISION_SERVICE = os.Getenv("DECISION_SERVICE")
var RSVP_SERVICE = os.Getenv("RSVP_SERVICE")

func LoadArborConfig() {
	security.AccessLogLocation = "log/access.log"
	security.ClientRegistryLocation = "clients.db"
	proxy.AccessControlPolicy = "*"
}
