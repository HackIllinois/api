package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/HackIllinois/api/common/config"
)

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

func (err *ApiError) Error() string {
	return err.Message
}

type ErrorLogEntry struct {
	ID    string
	Error interface{}
	Stack string
}

func LogError(id string, error_message interface{}) {
	log_entry := ErrorLogEntry{
		ID:    id,
		Error: error_message,
		Stack: string(debug.Stack()),
	}

	var error_log_message []byte
	var err error
	if config.DEBUG_MODE {
		error_log_message, err = json.MarshalIndent(log_entry, "", "    ")
	} else {
		error_log_message, err = json.Marshal(log_entry)
	}

	if err != nil {
		fmt.Printf("Failed to marshal error for id: %v\n", id)
		return
	}

	fmt.Printf("ERROR: %v\n", string(error_log_message))
}

// Writes the given error to the passed HTTP response
func WriteError(w http.ResponseWriter, r *http.Request, err ApiError) {
	LogError(r.Header.Get("HackIllinois-Identity"), err)

	// Strip the raw error string if we're not in debug mode
	if !config.DEBUG_MODE {
		err.RawError = ""
	}

	w.WriteHeader(err.Status)

	json.NewEncoder(w).Encode(err)
}
