package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var IS_PRODUCTION bool
var DEBUG_MODE bool

func init() {
	err := Initialize()

	if err != nil {
		panic(err)
	}
}

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	production, err := cfg_loader.Get("IS_PRODUCTION")

	if err != nil {
		return err
	}

	IS_PRODUCTION = (production == "true")

	debug_mode, err := cfg_loader.Get("DEBUG_MODE")

	if err != nil {
		return err
	}

	DEBUG_MODE = (debug_mode == "true")

	return nil
}
