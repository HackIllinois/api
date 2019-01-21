package service

import (
	"errors"
	"github.com/HackIllinois/api/services/auth/models"
	"net/url"
	"strings"
)

type OAuthProvider interface {
	GetAuthorizationRedirect(redirect_uri string) (string, error)
	Authorize(code string, redirect_uri string) error
	GetUserInfo() (*models.UserInfo, error)
	IsVerifiedUser() bool
}

/*
	Returns an OAuth provider struct for the requested provider
*/
func GetOAuthProvider(provider string) (OAuthProvider, error) {
	switch provider {
	case "github":
		return NewGitHubOAuth(), nil
	case "google":
		return NewGoogleOAuth(), nil
	case "linkedin":
		return NewLinkedInOAuth(), nil
	default:
		return nil, errors.New("Invalid provider")
	}
}

/*
	A helper function that takes a URL pointer and a map of query params->values, and modifies the URL's
	RawQuery property with the supplied query params.
*/
func ConstructURLQuery(u *url.URL, params map[string]string) {
	q := u.Query()

	for param, value := range params {
		q.Set(param, value)
	}

	u.RawQuery = q.Encode()
}

/*
	This function takes in the ingredients to a URL and outputs a string of them all together.
	It also checks for the appearance of "#" anywhere in the query params and throws an error if it is there.
    queryParams is an optional param. nil can be passed in if the url needs no query params.
*/
func ConstructSafeURL(scheme string, host string, path string, queryParams map[string]string) (string, error) {
	url := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}

	// Per the OAuth 2.0 RFC 6749, we need to disallow the `#` fragment character in the URL
	if queryParams != nil {
		for _, val := range queryParams {
			if strings.Contains(val, "#") {
				return url.String(), errors.New("`#` is an invalid character")
			}
		}

		ConstructURLQuery(&url, queryParams)
	}

	return url.String(), nil
}
