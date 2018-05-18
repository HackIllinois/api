package models

type DecisionHistory struct {
	ID      string     `json:"id"`
	Status  string     `json:"status"`
	Wave    int        `json:"wave"`
	History []Decision `json:"history"`
}
