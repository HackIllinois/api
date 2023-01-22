package config

import (
	"os"
	"strconv"

	"github.com/HackIllinois/api/common/configloader"
	"github.com/arbor-dev/arbor/proxy"
	"github.com/arbor-dev/arbor/security"
)

var GATEWAY_PORT uint16

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
var NOTIFICATIONS_SERVICE string
var PROJECT_SERVICE string
var PROFILE_SERVICE string

func Initialize() error {

	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	AUTH_SERVICE, err = cfg_loader.Get("AUTH_SERVICE")

	if err != nil {
		return err
	}

	USER_SERVICE, err = cfg_loader.Get("USER_SERVICE")

	if err != nil {
		return err
	}

	REGISTRATION_SERVICE, err = cfg_loader.Get("REGISTRATION_SERVICE")

	if err != nil {
		return err
	}

	DECISION_SERVICE, err = cfg_loader.Get("DECISION_SERVICE")

	if err != nil {
		return err
	}

	RSVP_SERVICE, err = cfg_loader.Get("RSVP_SERVICE")

	if err != nil {
		return err
	}

	CHECKIN_SERVICE, err = cfg_loader.Get("CHECKIN_SERVICE")

	if err != nil {
		return err
	}

	UPLOAD_SERVICE, err = cfg_loader.Get("UPLOAD_SERVICE")

	if err != nil {
		return err
	}

	MAIL_SERVICE, err = cfg_loader.Get("MAIL_SERVICE")

	if err != nil {
		return err
	}

	EVENT_SERVICE, err = cfg_loader.Get("EVENT_SERVICE")

	if err != nil {
		return err
	}

	STAT_SERVICE, err = cfg_loader.Get("STAT_SERVICE")

	if err != nil {
		return err
	}

	NOTIFICATIONS_SERVICE, err = cfg_loader.Get("NOTIFICATIONS_SERVICE")

	if err != nil {
		return err
	}

	PROJECT_SERVICE, err = cfg_loader.Get("PROJECT_SERVICE")

	if err != nil {
		return err
	}

	PROFILE_SERVICE, err = cfg_loader.Get("PROFILE_SERVICE")

	if err != nil {
		return err
	}

	port_str, err := cfg_loader.Get("GATEWAY_PORT")

	if err != nil {
		return err
	}

	port, err := strconv.ParseUint(port_str, 10, 16)

	if err != nil {
		return err
	}

	GATEWAY_PORT = uint16(port)

	return nil
}

func LoadArborConfig() {
	security.AccessLogLocation = "log/access.log"
	security.ClientRegistryLocation = "clients.db"
	proxy.AccessControlPolicy = "*"
}
