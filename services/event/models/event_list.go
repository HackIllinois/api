package models

type EventList[T Event] struct {
	Events []T `json:"events"`
}
