package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateSignedToken(secret []byte, data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(secret)
}

func ExtractFieldFromJWT(secret string, token string, field string) ([]string, error) {
	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
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
