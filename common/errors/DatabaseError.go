package errors

import "net/http"

// An error that occurs when a database operation (e.g. fetch / insert / update) doesn't work.
func DatabaseError(raw_error string, message string) ApiError {
	return ApiError{Status: http.StatusInternalServerError, Type: "DATABASE_ERROR", Message: message, RawError: raw_error}
}
