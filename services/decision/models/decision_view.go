package models

type DecisionView struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	ExpiresAt int64  `json:"expiresAt" `
}
