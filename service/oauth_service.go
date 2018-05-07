package service

import (
	"errors"
	"github.com/hackillinois/api-auth/config"
)

/*
	Return the oauth authoization url for the given provider
*/
func GetAuthorizeRedirect(provider string) (string, error) {
	switch provider {
	case "github":
		return "https://github.com/login/oauth/authorize?client_id=" + config.GITHUB_CLIENT_ID, nil
	default:
		return "", errors.New("Invalid provider")
	}
}

/*
	Gets the user's email from the specified oauth provider
*/
func GetEmail(oauth_token string, provider string) (string, error) {
	switch provider {
	case "github":
		return GetGithubEmail(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}

/*
	Converts an oauth code to an oauth token for the specified provider
*/
func GetOauthToken(code string, provider string) (string, error) {
	switch provider {
	case "github":
		return GetGithubOauthToken(code)
	default:
		return "", errors.New("Invalid provider")
	}
}

/*
	Gets the user's unique id from the specified oauth provider
*/
func GetUniqueId(oauth_token string, provider string) (string, error) {
	switch provider {
	case "github":
		return GetGithubUniqueId(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}
