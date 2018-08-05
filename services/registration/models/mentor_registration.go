package models

type MentorRegistration struct {
	ID        string `json:"id"          validate:"required"`
	FirstName string `json:"firstName"   validate:"required"`
	LastName  string `json:"lastName"    validate:"required"`
	Email     string `json:"email"       validate:"required,email"`
	ShirtSize string `json:"shirtSize"   validate:"required,oneof=S M L XL"`
	GitHub    string `json:"github"      validate:"required"`
	Linkedin  string `json:"linkedin"    validate:"required"`
	CreatedAt int64  `json:"createdAt"   validate:"required"`
	UpdatedAt int64  `json:"updatedAt"   validate:"required"`
}
