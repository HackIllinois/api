package models

type User struct {
	ID      string   `json:"id"`
	Devices []string `json:"devices"`
}
