package models

type DecisionHistory struct {
	Finalized bool       `json:"finalized"`
	ID        string     `json:"id"`
	Status    string     `json:"status"`
	Wave      int        `json:"wave"`
	Reviewer  string     `json:"reviewer"`
	Timestamp int64      `json:"timestamp"`
	ExpiresAt int64      `json:"expiresAt"`
	History   []Decision `json:"history"`
}
