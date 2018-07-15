package models

type MailOrder struct {
	IDs      []string `json:"ids"`
	Template string   `json:"template"`
}
