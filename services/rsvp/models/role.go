package models

type Role string

const (
	User Role = "User"
	Applicant Role = "Applicant"
	Attendee Role = "Attendee"
	Mentor Role = "Mentor"
	Admin Role = "Admin"
)
