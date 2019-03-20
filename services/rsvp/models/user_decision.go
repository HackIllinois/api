package models

type UserDecision struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Wave      int    `json:"wave"`
	Finalized bool   `json:"finalized"`
	Timestamp int64  `json:"timestamp"`
	ExpiresAt int64  `json:"expiresAt"`
}
