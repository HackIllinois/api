package models

type LinkedinUserInfo struct {
	ID        string `json:"id"`
	Username  string `json:"formattedName"`
	Email     string `json:"emailAddress"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
