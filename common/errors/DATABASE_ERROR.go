package errors

import "net/http"

// An error that occurs when a database operation (e.g. fetch / insert / update) doesn't work.
func DATABASE_ERROR(message string) APIError {
	return APIError{Status: http.StatusInternalServerError, Type: "DatabaseError", Message: message}
}
