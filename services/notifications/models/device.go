package models

type Device struct {
	UserID        string            `json:"userId"`
	DeviceToken   string            `json:"deviceToken"`
	DeviceArn     string            `json:"deviceArn"`
	Platform      string            `json:"platform"`
	Subscriptions map[string]string `json:"subscriptions"`
}
