package models

type DeviceRegistration struct {
	DeviceToken string `json:"deviceToken"`
	Platform    string `json:"platform"`
}
