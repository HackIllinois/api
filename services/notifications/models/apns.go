package models

type APNSPayload struct {
	Container APNSContainer    `json:"aps"`
	Data      PastNotification `json:"data"`
}

type APNSContainer struct {
	Alert APNSAlert `json:"alert"`
	Sound string    `json:"sound"`
}

type APNSAlert struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
