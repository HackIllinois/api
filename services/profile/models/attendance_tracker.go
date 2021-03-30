package models

type AttendanceTracker struct {
	ID     string   `json:"id"`
	Events []string `json:"events"`
}
