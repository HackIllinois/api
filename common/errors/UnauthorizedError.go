package errors

func UnauthorizedError(message string) APIError {
	return APIError{Status: 403, Title: "Invalid Authorization", Message: message}
}
