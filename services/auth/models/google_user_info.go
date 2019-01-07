package models

type GoogleUserInfo struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	FirstName  string `json:"given_name"`
	LastName   string `json:"family_name"`
	IsVerified bool   `json:"verified_email"`
}
