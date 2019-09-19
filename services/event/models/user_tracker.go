package models

type UserTracker struct {
	UserID string   `json:"userId"`
	Events map[string]int `json:"events"`
}
