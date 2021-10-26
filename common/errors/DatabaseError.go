package errors

import (
	"net/http"
)

const MONGO_NOT_FOUND_ERR = "Error: NOT_FOUND"

// An error that occurs when a database operation (e.g. fetch / insert / update) doesn't work.
func DatabaseError(raw_error string, message string) ApiError {
	http_status_error := http.StatusInternalServerError
	if raw_error == MONGO_NOT_FOUND_ERR {
		http_status_error = http.StatusNotFound
	}

	return ApiError{Status: http_status_error, Type: "DATABASE_ERROR", Message: message, RawError: raw_error}
}
