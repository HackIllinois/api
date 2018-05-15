package middleware

import (
	"github.com/HackIllinois/api-gateway/utils"
	"github.com/justinas/alice"
	"net/http"
)

func AuthMiddleware(authorized_roles []string) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			authorized, err := IsAuthorized(token, authorized_roles)
			if err != nil || !authorized {
				w.WriteHeader(403)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func IsAuthorized(token string, authorized_roles []string) (bool, error) {
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
