package config

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"

	"github.com/HackIllinois/api/common/configloader"
)

var IS_PRODUCTION bool
var DEBUG_MODE bool
var TOKEN_SECRET string

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

	TOKEN_SECRET, err = cfg_loader.Get("TOKEN_SECRET")

	if err != nil {
		return err
	}

	return nil
}

func GenerateIdFromSignedToken(signed_token string) (interface{}, error) {
	token, err := jwt.Parse(signed_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(TOKEN_SECRET), nil
	})

	if err != nil {
		return "", err
	}

	id := token.Claims.(jwt.MapClaims)["userId"]

	return id, err
}
