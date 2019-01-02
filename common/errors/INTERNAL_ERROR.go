package errors

import "net/http"

// Represents errors in the system, including failures in inter-service API calls.
// If the source of an error is unidentified, we fall back to this error.
func INTERNAL_ERROR(message string) APIError {
	return APIError{Status: http.StatusInternalServerError, Type: "InternalError", Message: message}
}
