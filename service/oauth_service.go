package service

import (
	"errors"
	"github.com/HackIllinois/api-auth/config"
	"strings"
)

/*
	Return the oauth authoization url for the given provider
*/
func GetAuthorizeRedirect(provider string) (string, error) {
	switch provider {
	case "github":
		return "https://github.com/login/oauth/authorize?client_id=" + config.GITHUB_CLIENT_ID, nil
	case "google":
		return "https://accounts.google.com/o/oauth2/v2/auth?client_id=" + config.GOOGLE_CLIENT_ID + "&redirect_uri=" + config.AUTH_REDIRECT_URI + "&scope=profile%20email&response_type=code", nil
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
	case "google":
		return GetGoogleEmail(oauth_token)
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
	case "google":
		return GetGoogleOauthToken(code)
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
	case "google":
		return GetGoogleUniqueId(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}

/*
	Gets the user's username from the specified oauth provider
*/
func GetUsername(oauth_token string, provider string) (string, error) {
	switch provider {
	case "github":
		return GetGithubUsername(oauth_token)
	case "google":
		return GetGoogleUsername(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}

/*
	Gets the user's first name from the specified oauth provider
*/
func GetFirstName(oauth_token string, provider string) (string, error) {
	const number_of_names int = 2
	const name_delimiter string = " "

	switch provider {
	case "github":
		name, err := GetGithubName(oauth_token)

		if err != nil {
			return "", err
		}

		split_name := strings.SplitAfterN(name, name_delimiter, number_of_names)

		return strings.TrimSpace(split_name[0]), nil
	case "google":
		return GetGoogleFirstName(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}

/*
	Gets the user's last name from the specified oauth provider
*/
func GetLastName(oauth_token string, provider string) (string, error) {
	const number_of_names int = 2
	const name_delimiter string = " "

	switch provider {
	case "github":
		name, err := GetGithubName(oauth_token)

		if err != nil {
			return "", err
		}

		split_name := strings.SplitAfterN(name, name_delimiter, number_of_names)

		// If there is only a single name, or if the name cannot be split.
		if len(split_name) < 2 {
			return "", nil
		} else {
			return strings.TrimSpace(split_name[1]), nil
		}
	case "google":
		return GetGoogleLastName(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}
