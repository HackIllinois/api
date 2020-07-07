package models

type AllRegistration struct {
	Attendee *RegistrationInfo `json:"attendee"`
	Mentor   *RegistrationInfo `json:"mentor"`
}
