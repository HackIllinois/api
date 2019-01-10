package errors

import "net/http"

// Used when the user attempts to perform an action that is not permitted as part of the flow.
// E.g. Attempting to check-in without having RSVPed.
func AttributeMismatchError(raw_error string, message string) ApiError {
	return ApiError{Status: http.StatusUnprocessableEntity, Type: "ATTRIBUTE_MISMATCH_ERROR", Message: message, RawError: raw_error}
}
