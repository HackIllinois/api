package errors

import "net/http"

// Represents errors in the system, including failures in inter-service API calls.
// If there are multiple possible sources of an error, we use this error.
func InternalError(raw_error string, message string) ApiError {
	return ApiError{Status: http.StatusInternalServerError, Type: "INTERNAL_ERROR", Message: message, RawError: raw_error}
}
