package models

type Blob struct {
	ID   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}
