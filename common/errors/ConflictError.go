package errors

import "net/http"

// Represents errors where items are not found.
// Can be used as general error status.
func ConflictError(raw_error string, message string) ApiError {
	return ApiError{Status: http.StatusConflict, Type: "CONFLICT_ERROR", Message: message, RawError: raw_error}
}
