package service

import (
	"github.com/ReflectionsProjections/api/services/auth/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

/*
	Generates a signed oauth token with the user's id, email, and roles embedded in the claims
*/
func MakeToken(id string, email string, roles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
		"id":    id,
		"email": email,
		"roles": roles,
	})

	signed_token, err := token.SignedString(config.TOKEN_SECRET)

	return signed_token, err
}
