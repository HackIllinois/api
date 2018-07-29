package models

type LinkedinOauthToken struct {
	Token    string `json:"access_token"`
	Lifetime int    `json:"expires_in"`
}
