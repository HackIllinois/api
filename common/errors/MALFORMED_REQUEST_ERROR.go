package errors

import "net/http"

// An error for when struct validation fails, or there are other issues with the payload to an endpoint.
func MALFORMED_REQUEST_ERROR(message string) APIError {
	return APIError{Status: http.StatusUnprocessableEntity, Type: "MalformedRequestError", Message: message}
}
