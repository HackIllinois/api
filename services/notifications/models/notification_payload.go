package models

type NotificationPayload struct {
	APNS    string `json:"APNS"`
	GCM     string `json:"GCM"`
	Default string `json:"default"`
}
