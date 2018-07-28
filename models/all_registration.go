package models

type AllRegistration struct {
	Attendee *UserRegistration   `json:"attendee"`
	Mentor   *MentorRegistration `json:"mentor"`
}
