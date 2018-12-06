package errors

/**
* Status is the HTTP error code to be sent to the client - should be set by constructor
* Type is the broad category - DatabaseError, AuthorizationError, InternalError
* Message provides additional details on the specific error that occurred.
 */
type APIError struct {
	Status  int    `json:"status"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
