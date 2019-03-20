package models

type Blob struct {
	ID   string      `json:"id"`
	Data interface{} `json:"data"`
}
