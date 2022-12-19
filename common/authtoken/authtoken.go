package authtoken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Role = string

const (
	AdminRole     = "Admin"
	StaffRole     = "Staff"
	MentorRole    = "Mentor"
	ApplicantRole = "Applicant"
	AttendeeRole  = "Attendee"
	UserRole      = "User"
	BlobstoreRole = "Blobstore"
)

func IsAuthorized(secret string, token string, authorized_roles []Role) (bool, error) {
	for _, role := range authorized_roles {
		is_authorized, err := HasRole(secret, token, role)

		if err != nil {
			return false, err
		}

		if is_authorized {
			return true, nil
		}
	}

	return false, nil
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

func HasRole(secret string, token string, required_role Role) (bool, error) {
	roles, err := ExtractFieldFromJWT(secret, token, "roles")

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
