package middleware

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/common/utils"

	common_config "github.com/HackIllinois/api/common/config"
)

func IdentificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		id, err := utils.ExtractFieldFromJWT(common_config.TOKEN_SECRET, token, "id")
		if err == nil {
			//Check if the user has the Admin role
			is_admin, err := authtoken.HasRole(common_config.TOKEN_SECRET, token, authtoken.AdminRole)
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
