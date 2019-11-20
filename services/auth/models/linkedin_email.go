package models

type LinkedinEmail struct {
	Elements []LinkedinEmailElement `json:"elements"`
}

type LinkedinEmailElement struct {
	Handle struct {
		Email string `json:"emailAddress"`
	} `json:"handle~"`
}
