package models

type EventTracker struct {
	EventID string   `json:"eventId"`
	Users   []string `json:"users"`
}
