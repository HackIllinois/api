package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"github.com/HackIllinois/api/common/errors"
	"os"
	"strings"
)

var IS_PRODUCTION bool
var DEBUG_MODE bool

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not load from config file."))
	}

	production, err := cfg_loader.Get("IS_PRODUCTION")

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not get variable IS_PRODUCTION from configloader."))
	}

	IS_PRODUCTION = (strings.ToLower(production) == "true")

	debug_mode, err := cfg_loader.Get("DEBUG_MODE")

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not get variable DEBUG_MODE from configloader."))
	}

	DEBUG_MODE = (strings.ToLower(debug_mode) == "true")
}
