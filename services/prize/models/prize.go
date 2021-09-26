package models

type Prize struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Value    int    `json:"value"`
	Quantity int    `json:"quantity"`
	ImageUrl string `json:"imgURL"`
}
