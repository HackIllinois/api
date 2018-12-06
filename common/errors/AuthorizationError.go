package errors

import "net/http"

// An error that occurs when an incoming request comes without a JWT token in the Authorization header.
func AuthorizationError(message string) APIError {
	return APIError{Status: http.StatusForbidden, Type: "AuthorizationError", Message: message}
}
