package config

import (
	"github.com/arbor-dev/arbor/proxy"
	"github.com/arbor-dev/arbor/security"
	"os"
	"strconv"
	"github.com/HackIllinois/api/common/configloader"
)

var GATEWAY_PORT uint16

var TOKEN_SECRET string

var AUTH_SERVICE string
var USER_SERVICE string
var REGISTRATION_SERVICE string
var DECISION_SERVICE string
var RSVP_SERVICE string
var CHECKIN_SERVICE string
var UPLOAD_SERVICE string
var MAIL_SERVICE string
var EVENT_SERVICE string
var STAT_SERVICE string

func init() {

	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	TOKEN_SECRET, err = cfg_loader.Get("TOKEN_SECRET")

	if err != nil {
		panic(err)
	}

	AUTH_SERVICE, err = cfg_loader.Get("AUTH_SERVICE")

	if err != nil {
		panic(err)
	}

	USER_SERVICE, err = cfg_loader.Get("USER_SERVICE")

	if err != nil {
		panic(err)
	}

	REGISTRATION_SERVICE, err = cfg_loader.Get("REGISTRATION_SERVICE")

	if err != nil {
		panic(err)
	}

	DECISION_SERVICE, err = cfg_loader.Get("DECISION_SERVICE")

	if err != nil {
		panic(err)
	}

	RSVP_SERVICE, err = cfg_loader.Get("RSVP_SERVICE")

	if err != nil {
		panic(err)
	}

	CHECKIN_SERVICE, err = cfg_loader.Get("CHECKIN_SERVICE")

	if err != nil {
		panic(err)
	}

	UPLOAD_SERVICE, err = cfg_loader.Get("UPLOAD_SERVICE")

	if err != nil {
		panic(err)
	}

	MAIL_SERVICE, err = cfg_loader.Get("MAIL_SERVICE")

	if err != nil {
		panic(err)
	}

	EVENT_SERVICE, err = cfg_loader.Get("EVENT_SERVICE")

	if err != nil {
		panic(err)
	}

	STAT_SERVICE, err = cfg_loader.Get("STAT_SERVICE")

	if err != nil {
		panic(err)
	}

	port_str, err := cfg_loader.Get("GATEWAY_PORT")

	if err != nil {
		panic(err)
	}

	port, err := strconv.ParseUint(port_str, 10, 16)

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
