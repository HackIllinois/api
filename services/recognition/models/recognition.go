package models

type Recognition struct {
	ID          string          `json:"id"                  validate:"required"`
	Name        string          `json:"name"                validate:"required"`
	Description string          `json:"description"         validate:"required"`
	StartTime   int64           `json:"startTime"           validate:"required"`
	EndTime     int64           `json:"endTime"             validate:"required"`
	Sponsor     string          `json:"sponsor"`
}
