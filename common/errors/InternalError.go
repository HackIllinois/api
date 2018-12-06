package errors

import "net/http"

// Represents errors in the system, including failures in inter-service API calls.
func InternalError(message string) APIError {
	return APIError{Status: http.StatusInternalServerError, Type: "InternalError", Message: message}
}
