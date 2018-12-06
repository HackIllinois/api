package errors

import "net/http"

// Used when the user attempts to perform an action that is not permitted as part of the flow.
// E.g. Attempting to check-in without having RSVPed.
func AttributeMismatchError(message string) APIError {
	return APIError{Status: http.StatusUnprocessableEntity, Type: "AttributeMismatchError", Message: message}
}
