package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/errors"
	"net/http"
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

	error_log_message, err := json.Marshal(log_entry)

	if err != nil {
		fmt.Printf("Failed to marshal error for id: %v\n", id)
		return
	}

	fmt.Printf("ERROR: %v\n", string(error_log_message))
}

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				w.Header().Set("Content-Type", "application/json")
				switch panic_error := rec.(type) {
				case errors.APIError:
					LogError(r.Header.Get("HackIllinois-Identity"), panic_error)
					w.WriteHeader(panic_error.Status)
					json.NewEncoder(w).Encode(panic_error)
				default:
					LogError(r.Header.Get("HackIllinois-Identity"), rec)
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(errors.APIError{Status: 500, Message: "Unknown Error"})
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
