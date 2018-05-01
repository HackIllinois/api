package middleware

import (
	"net/http"
	"../utils"
)

func IsAuthorized(r *http.Request, authorized_roles []string) (bool, error) {
	token := r.Header.Get("Authorization")

	for _, role := range authorized_roles {
		is_authorized, err := utils.HasRole(token, role)

		if err != nil {
			return false, err
		}

		if is_authorized {
			return true, nil
		}
	}

	return false, nil
}
