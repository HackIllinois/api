package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var NOTIFICATIONS_DB_HOST string
var NOTIFICATIONS_DB_NAME string

var NOTIFICATIONS_PORT string

var IS_PRODUCTION bool

var SNS_REGION string

var ANDROID_PLATFORM_ARN string
var IOS_PLATFORM_ARN string

var AUTH_SERVICE string

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	NOTIFICATIONS_DB_HOST, err = cfg_loader.Get("NOTIFICATIONS_DB_HOST")

	if err != nil {
		return err
	}

	NOTIFICATIONS_DB_NAME, err = cfg_loader.Get("NOTIFICATIONS_DB_NAME")

	if err != nil {
		return err
	}

	NOTIFICATIONS_PORT, err = cfg_loader.Get("NOTIFICATIONS_PORT")

	if err != nil {
		return err
	}

	SNS_REGION, err = cfg_loader.Get("SNS_REGION")

	if err != nil {
		return err
	}

	ANDROID_PLATFORM_ARN, err = cfg_loader.Get("ANDROID_PLATFORM_ARN")

	if err != nil {
		return err
	}

	IOS_PLATFORM_ARN, err = cfg_loader.Get("IOS_PLATFORM_ARN")

	if err != nil {
		return err
	}

	AUTH_SERVICE, err = cfg_loader.Get("AUTH_SERVICE")

	if err != nil {
		return err
	}

	production, err := cfg_loader.Get("IS_PRODUCTION")

	if err != nil {
		return err
	}

	IS_PRODUCTION = (production == "true")

	return nil
}
