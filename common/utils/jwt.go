package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateSignedToken(secret []byte, data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(secret)
}

func FetchIdFromSignedUserToken(secret string, signed_token string) (string, error) {
	token, err := jwt.Parse(signed_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	id, ok := token.Claims.(jwt.MapClaims)["userId"]
	if !ok {
		return "", fmt.Errorf("userId is nonexistent")
	}

	id_str, ok := id.(string)
	if !ok {
		return "", fmt.Errorf("Failed to cast id to string")
	}

	return id_str, nil
}

func ExtractFieldFromJWT(secret string, token_string string, field string) ([]string, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, fmt.Errorf("Expired token")
	}

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// We need to validate exp field exists - jwt will not do this on its own, as:
		// Valid validates time based claims "exp ...
		// if any of the above claims are not in the token, it will still be considered a valid claim.
		if claims["exp"] == nil {
			return nil, fmt.Errorf("Invalid token: 'exp' field missing")
		}

		_, ok := claims["exp"].(float64)
		if !ok {
			return nil, fmt.Errorf("Invalid token: 'exp' field malformed")
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
