package service

import (
	"time"

	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/golang-jwt/jwt/v4"
)

/*
Generates a signed oauth token with the user's id, email, and roles embedded in the claims
*/
func MakeToken(user_info *models.UserInfo, roles []string) (string, error) {
	return utils.GenerateSignedToken(config.TOKEN_SECRET, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 168).Unix(),
		"id":    user_info.ID,
		"email": user_info.Email,
		"roles": roles,
	})
}
