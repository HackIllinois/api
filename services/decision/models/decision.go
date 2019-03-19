package models

type Decision struct {
	Finalized bool   `json:"finalized" validate:"required|isdefault"`
	ID        string `json:"id"        validate:"required"`
	Status    string `json:"status"    validate:"required,oneof=PENDING REJECTED WAITLISTED ACCEPTED"`
	Wave      int    `json:"wave"      validate:""`
	Reviewer  string `json:"reviewer"  validate:"required"`
	Timestamp int64  `json:"timestamp" validate:"required"`
	ExpiresAt int64  `json:"expiresAt" validate:"required"`
}
