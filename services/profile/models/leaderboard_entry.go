package models

type LeaderboardEntry struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    int    `json:"points"`
}
