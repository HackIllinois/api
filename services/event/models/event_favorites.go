package models

type EventFavorites struct {
	ID     string   `json:"id"`
	Events []string `json:"events"`
}
