package models

type Topic struct {
	ID      string   `json:"id"`
	UserIDs []string `json:"userIds"`
}
