package models

type ProjectFavorites struct {
	ID       string   `json:"id"`
	Projects []string `json:"projects"`
}
