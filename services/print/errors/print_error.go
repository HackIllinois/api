package errors

import(
  "github.com/HackIllinois/api/common/errors"
)

func UnprocessableError(message string) APIError {
	return errors.APIError{Status: 400, Title: "Failed to publish print job", Message: message} // TODO reformat error body once APIError adds Type and Source fields
}
