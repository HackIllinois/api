package models

type GCMPayload struct {
	Notification GCMNotification `json:"notification"`
}

type GCMNotification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
