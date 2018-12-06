package errors

import "net/http"

// An error for when struct validation fails, or there are other issues with the payload to an endpoint.
func MalformedRequestError(message string) APIError {
	return APIError{Status: http.StatusUnprocessableEntity, Type: "MalformedRequestError", Message: message}
}
