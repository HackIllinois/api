package models

type ProfileFavorites struct {
	ID       string   `json:"id"`
	Profiles []string `json:"profiles"`
}
