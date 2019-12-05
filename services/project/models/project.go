package models

type Project struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Mentors  []string        `json:"mentors"`
	Location ProjectLocation `json:"location"`
	Tags     []string        `json:"tags"`
	Code     string          `json:"code"`
}

type ProjectLocation struct {
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
