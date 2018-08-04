package models

type UserTracker struct {
	UserID string   `json:"userId"`
	Events []string `json:"events"`
}
