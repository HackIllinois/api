package models

type APNSPayload struct {
	Container APNSContainer    `json:"aps"`
	Data      PastNotification `json:"data"`
}
