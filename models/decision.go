package models

type Decision struct {
	ID       string `json:"id"    validate:"required"`
	Status   string `json:"status" validate:"required,oneof=PENDING REJECTED WAITLISTED ACCEPTED"`
	Wave     int    `json:"wave"  validate:""`
}
