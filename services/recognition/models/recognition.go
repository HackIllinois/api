package models

type Recognition struct {
	ID          string          `json:"id"                  validate:"required"`
	Name        string          `json:"name"                validate:"required"`
	Description string          `json:"description"         validate:"required"`
	Presenter   string          `json:"presenter"           validate:"required"`
	EventID     string          `json:"eventId"             validate:"required"`
	Recepients  []Recepient     `json:"recipients"          validate:"required"`
	Tags        []string        `json:"tags"`
}

type Recepient struct {
	Type        string `json:"type" validate:"required,oneof=ALL INDIVIDUAL"`
	UserID      string `json:"userId"`
}
