package middleware

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	common_config "github.com/HackIllinois/api/common/config"
	"github.com/justinas/alice"
)

func AuthMiddleware(authorized_roles []authtoken.Role) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			authorized, err := authtoken.IsAuthorized(common_config.TOKEN_SECRET, token, authorized_roles)
			if err != nil || !authorized {
				w.WriteHeader(403)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
