package models

type NotificationPayload struct {
	APNS        string `json:"APNS"`
	APNSSandbox string `json:"APNS_SANDBOX"`
	GCM         string `json:"GCM"`
	Default     string `json:"default"`
}
