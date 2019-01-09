package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/errors"
	"net/http"
	"os"
	"strings"
)

type ErrorLogEntry struct {
	ID    string
	Error interface{}
}

func LogError(id string, error_message interface{}) {
	log_entry := ErrorLogEntry{
		ID:    id,
		Error: error_message,
	}

	error_log_message, err := json.MarshalIndent(log_entry, "", "    ")

	if err != nil {
		fmt.Printf("Failed to marshal error for id: %v\n", id)
		return
	}

	fmt.Printf("ERROR: %v\n", string(error_log_message))
}

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if panic_object := recover(); panic_object != nil {
				w.Header().Set("Content-Type", "application/json")
				switch panic_error := panic_object.(type) {
				case errors.ApiError:
					handleApiError(panic_error, w, r)
				default:
					handleUnknownError(panic_error, w, r)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func handleApiError(err errors.ApiError, w http.ResponseWriter, r *http.Request) {
	LogError(r.Header.Get("HackIllinois-Identity"), err)

	w.WriteHeader(err.Status)

	mode, is_mode_set := os.LookupEnv("DEBUG_MODE")

	// Strip the raw error string if we're not in debug mode
	if !is_mode_set || (is_mode_set && strings.ToLower(mode) != "true") {
		err.RawError = ""
	}

	json.NewEncoder(w).Encode(err)
}

func handleUnknownError(err interface{}, w http.ResponseWriter, r *http.Request) {
	LogError(r.Header.Get("HackIllinois-Identity"), err)

	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(errors.ApiError{Status: http.StatusInternalServerError, Type: "UNKNOWN_ERROR", Message: err.(string), RawError: ""})
}
