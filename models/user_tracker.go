package models

type UserTracker struct {
	ID     string   `json:"id"`
	Events []string `json:"events"`
}
