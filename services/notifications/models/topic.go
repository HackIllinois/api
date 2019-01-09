package models

type Topic struct {
	Name    string   `json:"name"`
	Arn     string   `json:"arn"`
	UserIDs []string `json:"userIds"`
}
