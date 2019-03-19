package utils

import (
	"errors"
)

func ContainsString(slice []string, str string) bool {
	for _, value := range slice {
		if value == str {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, str string) ([]string, error) {
	for i, value := range slice {
		if value == str {
			return append(slice[:i], slice[i+1:]...), nil
		}
	}
	return nil, errors.New("Value to remove not found")
}
