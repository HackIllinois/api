package models

type EventCode struct {
	CodeID     string `json:"codeID"`
	EventID    string `json:"eventID"`
	IsVirtual  bool   `json:"isVirtual"`
	Expiration int64  `json:"expiration"`
}
