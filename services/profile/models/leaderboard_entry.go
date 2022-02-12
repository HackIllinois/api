package models

type LeaderboardEntry struct {
	ID      string `json:"id"`
	Points  int    `json:"points"`
	Discord string `json:"discord"`
}
