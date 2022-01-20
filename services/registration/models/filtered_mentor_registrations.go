package models

import (
	"encoding/json"
)

type FilteredMentorRegistrations struct {
	Registrations []MentorRegistration `json:"registrations"`
}

func (original FilteredMentorRegistrations) MarshalJSON() ([]byte, error) {
	type Alias FilteredMentorRegistrations

	modified := struct {
		Alias
	}{
		Alias: (Alias)(original),
	}

	if modified.Registrations == nil {
		modified.Registrations = make([]MentorRegistration, 0)
	}
	return json.Marshal(modified)
}
