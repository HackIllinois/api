package models

type GoogleOauthToken struct {
	Token    string `json:"access_token"`
	Type     string `json:"token_type"`
	Lifetime int    `json:"expires_in"`
	IDToken  string `json:"id_token"`
}
