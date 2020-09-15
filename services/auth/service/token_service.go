package service

import (
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

/*
	Generates a signed oauth token with the user's id, email, and roles embedded in the claims
*/
func MakeToken(user_info *models.UserInfo, roles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 168).Unix(),
		"id":    user_info.ID,
		"email": user_info.Email,
		"roles": roles,
	})

	signed_token, err := token.SignedString(config.TOKEN_SECRET)

	return signed_token, err
}
