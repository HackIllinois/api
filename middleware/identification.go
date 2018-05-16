package middleware

import (
	"github.com/HackIllinois/api-gateway/utils"
	"net/http"
)

func IdentificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		id, err := utils.ExtractFieldFromJWT(token, "id")

		if err == nil {
			r.Header.Set("HackIllinois-Identity", id[0])
		} else {
			r.Header.Set("HackIllinois-Identity", "")
		}

		next.ServeHTTP(w, r)
	})
}
