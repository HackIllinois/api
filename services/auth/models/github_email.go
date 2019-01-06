package models

type GithubEmail struct {
	Email      string `json:"email"`
	IsPrimary  bool   `json:"primary"`
	IsVerified bool   `json:"verified"`
}
