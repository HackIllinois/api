package models

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Mentors     []string `json:"mentors"`
	Room        string   `json:"room"`
	Tags        []string `json:"tags"`
	Number      int64    `json:"number"`
}
