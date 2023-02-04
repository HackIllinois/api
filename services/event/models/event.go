package models

type Event interface {
	EventPublic | EventDB
}

type EventPublic struct {
	ID          string          `json:"id"          validate:"required"`
	Name        string          `json:"name"        validate:"required"`
	Description string          `json:"description" validate:"required"`
	StartTime   int64           `json:"startTime"   validate:"required_if=IsAsync False"`
	EndTime     int64           `json:"endTime"     validate:"required_if=IsAsync False"`
	Locations   []EventLocation `json:"locations"   validate:"required,dive,required"`
	Sponsor     string          `json:"sponsor"`
	EventType   string          `json:"eventType"   validate:"required,oneof=MEAL SPEAKER WORKSHOP MINIEVENT QNA OTHER"`
	Points      int             `json:"points"`
	IsAsync     bool            `json:"isAsync"`
}

// Struct to encapsulate hidden fields
type EventDB struct {
	EventPublic           `     bson:",inline"`
	IsPrivate             bool `               json:"isPrivate"`
	DisplayOnStaffCheckin bool `               json:"displayOnStaffCheckin"`
}

type EventLocation struct {
	Description string   `json:"description" validate:"required"`
	Tags        []string `json:"tags"        validate:"required"`
	Latitude    float64  `json:"latitude"    validate:"required"`
	Longitude   float64  `json:"longitude"   validate:"required"`
}
