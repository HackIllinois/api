package models

type NotificationOrder struct {
	ID         string `json:"id"`
	Recipients int    `json:"recipients"`
	Success    int    `json:"success"`
	Failure    int    `json:"failure"`
	Time       int64  `json:"time"`
}
