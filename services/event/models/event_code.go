package models

type EventCode struct {
	CodeID     string `json:"codeID"     validate:"required"`
	EventID    string `json:"eventID"    validate:"required"`
	IsVirtual  bool   `json:"isVirtual"  validate:"required"`
	Expiration int64  `json:"expiration" validate:"required"`
}
