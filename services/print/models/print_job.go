package models

type PrintJob struct {
	ID       string        `json:"id" validate:"required"`
	Location PrintLocation `json:"location" validate:"required"`
}
