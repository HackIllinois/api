package errors

import "net/http"

// An error for when struct validation fails, or there are other issues with the payload to an endpoint.
func MalformedRequestError(raw_error string, message string) ApiError {
	return ApiError{Status: http.StatusUnprocessableEntity, Type: "MALFORMED_REQUEST_ERROR", Message: message, RawError: raw_error}
}
