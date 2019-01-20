package config

import (
	"github.com/HackIllinois/api/common/configloader"
	"os"
)

var UPLOAD_DB_HOST string
var UPLOAD_DB_NAME string

var UPLOAD_PORT string

var S3_REGION string
var S3_BUCKET string
var IS_PRODUCTION bool

func Initialize() error {
	cfg_loader, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		return err
	}

	UPLOAD_DB_HOST, err = cfg_loader.Get("UPLOAD_DB_HOST")

	if err != nil {
		return err
	}

	UPLOAD_DB_NAME, err = cfg_loader.Get("UPLOAD_DB_NAME")

	if err != nil {
		return err
	}

	UPLOAD_PORT, err = cfg_loader.Get("UPLOAD_PORT")

	if err != nil {
		return err
	}

	S3_REGION, err = cfg_loader.Get("S3_REGION")

	if err != nil {
		return err
	}

	S3_BUCKET, err = cfg_loader.Get("S3_BUCKET")

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
