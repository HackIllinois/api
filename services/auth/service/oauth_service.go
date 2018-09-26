package service

import (
	"errors"
	"github.com/HackIllinois/api/services/auth/config"
	"net/url"
	"strings"
)

/*
	Return the oauth authorization url for the given provider
*/
func GetAuthorizeRedirect(provider string, redirect_uri string) (string, error) {
	switch provider {
	case "github":
		return ConstructSafeURL("https", "github.com", "login/oauth/authorize",
			map[string]string{
				"client_id":    config.GITHUB_CLIENT_ID,
				"redirect_uri": redirect_uri,
			})
	case "google":
		return ConstructSafeURL("https", "accounts.google.com", "o/oauth2/v2/auth",
			map[string]string{
				"client_id":     config.GOOGLE_CLIENT_ID,
				"scope":         "profile email",
				"response_type": "code",
				"redirect_uri":  redirect_uri,
			})
	case "linkedin":
		return ConstructSafeURL("https", "www.linkedin.com", "oauth2/v2/authorization",
			map[string]string{
				"client_id":     config.LINKEDIN_CLIENT_ID,
				"scope":         "r_basicprofile r_emailaddress",
				"response_type": "code",
				"redirect_uri":  redirect_uri,
			})
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
	case "linkedin":
		return GetLinkedinEmail(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}

/*
	Converts an oauth code to an oauth token for the specified provider
*/
func GetOauthToken(code string, provider string, redirect_uri string) (string, error) {
	switch provider {
	case "github":
		return GetGithubOauthToken(code)
	case "google":
		return GetGoogleOauthToken(code, redirect_uri)
	case "linkedin":
		return GetLinkedinOauthToken(code, redirect_uri)
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
	case "linkedin":
		return GetLinkedinUniqueId(oauth_token)
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
	case "linkedin":
		return GetLinkedinUsername(oauth_token)
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
	case "linkedin":
		return GetLinkedinFirstName(oauth_token)
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
	case "linkedin":
		return GetLinkedinLastName(oauth_token)
	default:
		return "", errors.New("Invalid provider")
	}
}

/*
	A helper function that takes a URL pointer and a map of query params->values, and modifies the URL's
	RawQuery property with the supplied query params.
*/
func constructURLQuery(u *url.URL, params map[string]string) {
	q := u.Query()

	for param, value := range params {
		q.Set(param, value)
	}

	u.RawQuery = q.Encode()
}

/*
	This function takes in the ingredients to a URl and outputs a string of them all together.
	It also checks for the apperance of "#" anywhere before the last query parameter, and returns an error if so.
*/
func ConstructSafeURL(scheme string, host string, path string, queryParams map[string]string) (string, error) {
	url := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}

    // Per the OAuth 2.0 RFC 6749, we need to disallow the `#` fragment character in the URL
	valid := true
	for _, val := range queryParams {
		if strings.Contains(val, "#") {
			valid = false
			break
		}
	}

	constructURLQuery(&url, queryParams)

	if valid {
		return url.String(), nil
	}

	return url.String(), errors.New("`#` is an invalid character")
}
