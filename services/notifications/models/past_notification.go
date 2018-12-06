package models

type PastNotification struct {
	TopicName string `json:"topicName"`
	Body      string `json:"body"`
	Title     string `json:"title"`
	Time      int64  `json:"time"`
}
