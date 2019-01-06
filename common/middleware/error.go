package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/errors"
	"net/http"
	"os"
	"strings"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if panic_object := recover(); panic_object != nil {
				w.Header().Set("Content-Type", "application/json")
				switch panic_object.(type) {
				case errors.ApiError:
					handleApiError(panic_object.(errors.ApiError), w)
				default:
					handleUnknownError(panic_object, w)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func handleApiError(err errors.ApiError, w http.ResponseWriter) {
	w.WriteHeader(err.Status)

	mode, is_mode_set := os.LookupEnv("DEBUG_MODE")

	// Strip the raw error string if we're not in debug mode
	if !is_mode_set || (is_mode_set && strings.ToLower(mode) != "true") {
		err.RawError = ""
	}

	json.NewEncoder(w).Encode(err)
}

func handleUnknownError(err interface{}, w http.ResponseWriter) {
	fmt.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errors.ApiError{Status: http.StatusInternalServerError, Type: "UNKNOWN_ERROR", Message: err.(string), RawError: ""})
}
