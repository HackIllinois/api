package middleware

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/justinas/alice"
)

func AuthMiddleware(authorized_roles []authtoken.Role) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			authorized, err := authtoken.IsAuthorized(config.TOKEN_SECRET, token, authorized_roles)
			if err != nil || !authorized {
				w.WriteHeader(403)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
