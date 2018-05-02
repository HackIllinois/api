package main

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

var secret []byte

func init() {
	secret = []byte("secret_string")
}

func MakeToken(id int, email string, roles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"id": id,
		"email": email,
		"roles": roles,
	})

	signed_token, err := token.SignedString(secret)

	return signed_token, err
}
