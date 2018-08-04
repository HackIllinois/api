package utils

import (
	"fmt"
	"github.com/HackIllinois/api/gateway/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var secret []byte

func init() {
	secret = []byte(config.TOKEN_SECRET)
}

func ExtractFieldFromJWT(token string, field string) ([]string, error) {
	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Invalid token")
	}

	if claims, ok := jwt_token.Claims.(jwt.MapClaims); ok && jwt_token.Valid {
		if int64(claims["exp"].(float64)) < time.Now().Unix() {
			return nil, fmt.Errorf("Expired token")
		}

		var data []string
		switch elem := claims[field].(type) {
		case []interface{}:
			for _, item := range elem {
				data = append(data, item.(string))
			}
		case interface{}:
			data = append(data, elem.(string))
		}
		return data, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

func HasRole(token string, required_role string) (bool, error) {
	roles, err := ExtractFieldFromJWT(token, "roles")

	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role == required_role {
			return true, nil
		}
	}

	return false, nil
}
