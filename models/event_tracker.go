package models

type EventTracker struct {
	EventName string   `json:"eventName"`
	Users     []string `json:"users"`
}
