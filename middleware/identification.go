package middleware

import (
	"net/http"

	"github.com/HackIllinois/api-gateway/utils"
)

func IdentificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		id, err := utils.ExtractFieldFromJWT(token, "id")

		if err == nil {
			//Check if the user has role "Admin"
			is_admin, err := utils.HasRole(token, "Admin")

			if err == nil {
				if is_admin {
					impersonation_id := r.Header.Get("HackIllinois-Impersonation")

					// Check if an impersonation ID is specified
					// According to https://golang.org/src/net/http/header.go:37, Get returns an empty string if there is no value associated with a given header.

					if impersonation_id == "" {
						r.Header.Set("HackIllinois-Identity", id[0])
					} else {
						// Sets ID to the one specified by a user with role "Admin"
						r.Header.Set("HackIllinois-Identity", impersonation_id)
					}
				} else {
					r.Header.Set("HackIllinois-Identity", id[0])
				}
			} else {
				// Unable to determine the user's roles
				r.Header.Set("HackIllinois-Identity", id[0])
			}
		} else {
			// Cannot retrieve user's ID
			r.Header.Set("HackIllinois-Identity", "")
		}

		next.ServeHTTP(w, r)
	})
}
