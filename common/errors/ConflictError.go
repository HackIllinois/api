package errors

import "net/http"

// An error that occurs when an incoming request tries to create a resource that already exists.
func ConflictError(raw_error string, message string) ApiError {
	return ApiError{Status: http.StatusConflict, Type: "CONFLICT_ERROR", Message: message, RawError: raw_error}
}
