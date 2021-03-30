package models

type EventCode struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Expiration int64  `json:"expiration"`
}
