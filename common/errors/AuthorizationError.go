package errors

import "net/http"

// An error that occurs when an incoming request comes without a JWT token in the Authorization header.
func AuthorizationError(raw_error string, message string) ApiError {
	return ApiError{Status: http.StatusForbidden, Type: "AUTHORIZATION_ERROR", Message: message, RawError: raw_error}
}
