package models

type MailList struct {
	ID      string   `json:"id"`
	UserIDs []string `json:"userIds"`
}
