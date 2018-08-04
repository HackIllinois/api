package errors

func UnprocessableError(message string) APIError {
	return APIError{Status: 400, Title: "Unprocessable Request", Message: message}
}
