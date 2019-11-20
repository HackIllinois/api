package models

type NotificationOrder struct {
	ID            string `json:"id"`
	NumRecipients int    `json:"numRecipients"`
	Success       int    `json:"success"`
	Failure       int    `json:"failure"`
	Time          int64  `json:"time"`
}
