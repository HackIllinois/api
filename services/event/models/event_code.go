package models

type EventCode struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	IsVirtual  bool   `json:"isVirtual"`
	Expiration int64  `json:"expiration"`
}
