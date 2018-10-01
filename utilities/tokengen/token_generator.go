package main

import (
	"github.com/HackIllinois/api/utilities/tokengen/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*
	Generates a jwt signed with the given secret containing the provided experation, id, email, and roles in the claims
*/
func MakeToken(id string, exp int64, email string, roles []models.Role, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   exp,
		"id":    id,
		"email": email,
		"roles": roles,
	})

	signed_token, err := token.SignedString(secret)

	return signed_token, err
}
