package models

type GithubUserInfo struct {
	Username string `json:"login"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
}
