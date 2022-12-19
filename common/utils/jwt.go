package utils

import "github.com/golang-jwt/jwt/v4"

func GenerateSignedToken(secret []byte, data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(secret)
}
