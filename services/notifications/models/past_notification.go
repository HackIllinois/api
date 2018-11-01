package models

type PastNotification struct {
	TopicName string `json:"topicName"`
	Message   string `json:"message"`
	Time      int64  `json:"time"`
}
