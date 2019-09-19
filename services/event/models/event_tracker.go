package models

type EventTracker struct {
	EventID string   `json:"eventId"`
	Users   map[string]int `json:"users"`
}
