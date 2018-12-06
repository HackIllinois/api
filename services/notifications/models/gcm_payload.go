package models

type GCMPayload struct {
	Notification GCMNotification `json:"notification"`
}
