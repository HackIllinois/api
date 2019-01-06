package errors

/**
* Status - the HTTP error code to be sent to the client - should be set by constructor
* Type - the broad category - e.g. DatabaseError, AuthorizationError, InternalError
* Message - provides additional details on the specific error that occurred.
* RawError - the raw error (stringified) that caused the panic. It is only included in the response
* to the client, if the config variable DEBUG_MODE is set to true. In other cases, the
* field is set to the empty string, which causes its omission when encoded to JSON.
**/
type ApiError struct {
	Status   int    `json:"status"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	RawError string `json:"raw_error,omitempty"`
}
