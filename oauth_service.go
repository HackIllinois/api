package main

import (
	"errors"
	"./config"
)

func GetAuthorizeRedirect(provider string) (string, error) {
	switch provider {
	case "github":
		return "https://github.com/login/oauth/authorize?client_id=" + config.GITHUB_CLIENT_ID, nil
	default:
		return "", errors.New("Invalid provider")
	}
}

func GetEmail(oauth_token string, provider string) (string, error) {
	switch provider {
	case "github":
		return GetGithubEmail(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}

func GetOauthToken(code string, provider string) (string, error) {
	switch provider {
	case "github":
		return GetGithubOauthToken(code)
	default:
		return "", errors.New("Invalid provider")
	}
}

func GetUniqueId(oauth_token string, provider string) (string, error) {
	switch provider {
	case "github":
		return GetGithubUniqueId(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}
