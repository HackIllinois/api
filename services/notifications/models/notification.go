package models

type Notification struct {
	ID    string `json:"id"`
	Topic string `json:"topic"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Time  int64  `json:"time"`
}
