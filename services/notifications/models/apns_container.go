package models

type APNSContainer struct {
	Alert APNSAlert `json:"alert"`
	Sound string    `json:"sound"`
}
