package models

import (
	"encoding/json"
)

type FilteredUserRegistrations struct {
	Registrations []UserRegistration `json:"registrations"`
}

func (original FilteredUserRegistrations) MarshalJSON() ([]byte, error) {
	type Alias FilteredUserRegistrations

	modified := struct {
		Alias
	}{
		Alias: (Alias)(original),
	}

	if modified.Registrations == nil {
		modified.Registrations = make([]UserRegistration, 0)
	}
	return json.Marshal(modified)
}
