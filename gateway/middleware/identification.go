package middleware

import (
	"net/http"

	"github.com/ReflectionsProjections/api/gateway/utils"
)

func IdentificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		id, err := utils.ExtractFieldFromJWT(token, "id")
		if err == nil {
			//Check if the user has role "Admin"
			is_admin, err := utils.HasRole(token, "Admin")
			if err == nil && is_admin {
				impersonation_id := r.Header.Get("HackIllinois-Impersonation")
				// Check if an impersonation ID is specified
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
			// Cannot retrieve user's ID
			r.Header.Set("HackIllinois-Identity", "")
		}
		next.ServeHTTP(w, r)
	})
}
