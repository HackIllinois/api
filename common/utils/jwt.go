package utils

import "github.com/dgrijalva/jwt-go"

func GenerateSignedToken(secret []byte, data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(secret)
}
