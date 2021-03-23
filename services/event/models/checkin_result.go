package models

type CheckinResult struct {
	NewPoints   int    `json:"newPoints"`
	TotalPoints int    `json:"totalPoints"`
	Status      string `json:"status"`
}
