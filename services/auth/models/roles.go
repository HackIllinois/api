package models

type Role = string

const (
	AdminRole     = "Admin"
	StaffRole     = "Staff"
	MentorRole    = "Mentor"
	ApplicantRole = "Applicant"
	AttendeeRole  = "Attendee"
	UserRole      = "User"
	SponsorRole   = "Sponsor"
	BlobstoreRole = "Blobstore"
)

var Roles []Role = []Role{AdminRole, StaffRole, MentorRole, ApplicantRole, AttendeeRole, UserRole, SponsorRole, BlobstoreRole}
