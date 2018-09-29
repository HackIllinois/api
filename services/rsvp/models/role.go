package models

type Role int

const (
	User Role = iota
	Applicant
	Attendee
	Mentor
	Admin
)
