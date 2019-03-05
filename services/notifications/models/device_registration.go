package models

type DeviceRegistration struct {
	Token    string `json:"token"`
	Platform string `json:"platform"`
}
