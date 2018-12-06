package models

type TopicPublic struct {
	Name    string   `json:"name"`
	UserIDs []string `json:"userIds"`
}
