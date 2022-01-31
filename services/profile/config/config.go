package config

import (
	"os"

	"github.com/HackIllinois/api/common/configloader"
	"github.com/HackIllinois/api/services/profile/models"
)

var PROFILE_DB_HOST string
var PROFILE_DB_NAME string

var PROFILE_PORT string

var TIER_THRESHOLDS []models.TierThreshold

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	PROFILE_DB_HOST, err = cfg_loader.Get("PROFILE_DB_HOST")

	if err != nil {
		return err
	}

	PROFILE_DB_NAME, err = cfg_loader.Get("PROFILE_DB_NAME")

	if err != nil {
		return err
	}

	PROFILE_PORT, err = cfg_loader.Get("PROFILE_PORT")

	if err != nil {
		return err
	}

	err = cfg_loader.ParseInto("TIER_THRESHOLDS", &TIER_THRESHOLDS)

	if err != nil {
		return err
	}

	return nil
}
