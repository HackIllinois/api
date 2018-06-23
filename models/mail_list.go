package models

type MailList struct {
	ListID  string   `json:"listId"`
	UserIDs []string `json:"userIds"`
}
