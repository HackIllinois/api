package main

import (
	"time"
	"./config"
	jwt "github.com/dgrijalva/jwt-go"
)

func MakeToken(id int, email string, roles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"id": id,
		"email": email,
		"roles": roles,
	})

	signed_token, err := token.SignedString(config.TOKEN_SECRET)

	return signed_token, err
}
