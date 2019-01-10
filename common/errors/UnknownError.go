package errors

import "net/http"

// Represents errors in the system whose cause is unidentified.
func UnknownError(raw_error string, message string) ApiError {
	return ApiError{Status: http.StatusInternalServerError, Type: "UNKNOWN_ERROR", Message: message, RawError: raw_error}
}
