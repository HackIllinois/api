package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var IS_PRODUCTION bool

func init() {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		panic(err)
	}

	production, err := cfg_loader.Get("IS_PRODUCTION")

	if err != nil {
		panic(err)
	}

	IS_PRODUCTION = (production == "true")
}
