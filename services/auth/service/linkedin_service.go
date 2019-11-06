package service

import (
	"errors"

	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/levigross/grequests"
)

type LinkedInOAuthProvider struct {
	token          string
	isVerifiedUser bool
}

func NewLinkedInOAuth() *LinkedInOAuthProvider {
	return &LinkedInOAuthProvider{
		token:          "",
		isVerifiedUser: false,
	}
}

/*
	Returns the url to redirects to for OAuth authorization
*/
func (provider *LinkedInOAuthProvider) GetAuthorizationRedirect(redirect_uri string) (string, error) {
	return ConstructSafeURL("https", "www.linkedin.com", "oauth/v2/authorization",
		map[string]string{
			"client_id":     config.LINKEDIN_CLIENT_ID,
			"scope":         "r_liteprofile r_emailaddress", // changed from r_basicprofile to r_liteprofile
			"response_type": "code",
			"redirect_uri":  redirect_uri,
		},
	)
}

/*
	Exchanges an OAuth code for an OAuth token
*/
func (provider *LinkedInOAuthProvider) Authorize(code string, redirect_uri string) error {
	request, err := grequests.Post("https://www.linkedin.com/oauth/v2/accessToken", &grequests.RequestOptions{
		Data: map[string]string{
			"client_id":     config.LINKEDIN_CLIENT_ID,
			"client_secret": config.LINKEDIN_CLIENT_SECRET,
			"code":          code,
			"redirect_uri":  redirect_uri,
			"grant_type":    "authorization_code",
		},
		Headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/x-www-form-urlencoded",
			"x-li-format":  "json",
		},
	})

	if err != nil {
		return err
	}

	var oauth_token models.LinkedinOauthToken
	err = request.JSON(&oauth_token)

	if err != nil {
		return err
	}

	if oauth_token.Token == "" {
		return errors.New("Invalid OAuth code.")
	}

	provider.token = oauth_token.Token

	return nil
}

/*
	Retrieves user info from the OAuth provider
*/
func (provider *LinkedInOAuthProvider) GetUserInfo() (*models.UserInfo, error) {
	var user_info models.UserInfo

	// request, err := grequests.Get("https://api.linkedin.com/v1/people/~:(id,formatted-name,email-address,first-name,last-name)", &grequests.RequestOptions{
	// 	Headers: map[string]string{
	// 		"Authorization": "Bearer " + provider.token,
	// 		"Content-Type":  "application/json",
	// 		"x-li-format":   "json",
	// 	},
	// })

	// ISSUE: The formatted-name field is no longer available
	request, err := grequests.Get("https://api.linkedin.com/v2/me?projection=(id,firstName?fields=id,bar,baz,lastName)", &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": "Bearer " + provider.token,
			"Content-Type":  "application/json",
			"x-li-format":   "json",
		},
	})

	if err != nil {
		return nil, err
	}

	var linkedin_user_info models.LinkedinUserInfo
	err = request.JSON(&linkedin_user_info)

	if err != nil {
		return nil, err
	}

	if linkedin_user_info.ID == "" {
		return nil, errors.New("Invalid OAuth token.")
	}

	user_info.ID = "linkedin" + linkedin_user_info.ID

	preferred_country := linkedin_user_info.FirstName.PreferredLocale.Country
	preferred_language := linkedin_user_info.FirstName.PreferredLocale.Language
	if preferred_country != "" && preferred_language != "" {
		// preferred local is provided
		preferred_locale := preferred_language + "_" + preferred_country
		user_info.FirstName = linkedin_user_info.FirstName.Localized[preferred_locale]
		user_info.LastName = linkedin_user_info.LastName.Localized[preferred_locale]
		user_info.Username = linkedin_user_info.FirstName.Localized[preferred_locale] + " " + linkedin_user_info.LastName.Localized[preferred_locale]
	} else {
		// preferred local is not provided, try getting en_US first, if failed, pick an arbitrary one
		if linkedin_user_info.FirstName.Localized["en_US"] != "" && linkedin_user_info.LastName.Localized["en_US"] != "" {
			user_info.FirstName = linkedin_user_info.FirstName.Localized["en_US"]
			user_info.LastName = linkedin_user_info.LastName.Localized["en_US"]
			user_info.Username = linkedin_user_info.FirstName.Localized["en_US"] + "_" + linkedin_user_info.LastName.Localized["en_US"]
		} else {
			for locale, _ := range linkedin_user_info.FirstName.Localized {
				arbitrary_locale := locale
				user_info.FirstName = linkedin_user_info.FirstName.Localized[arbitrary_locale]
				user_info.LastName = linkedin_user_info.LastName.Localized[arbitrary_locale]
				user_info.Username = linkedin_user_info.FirstName.Localized[arbitrary_locale] + "_" + linkedin_user_info.LastName.Localized[arbitrary_locale]
				break
			}
		}
	}

	// In API v2, a seperate request is required to retrieve the email address
	request, err = grequests.Get("https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))", &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": "Bearer " + provider.token,
			"Content-Type":  "application/json",
			"x-li-format":   "json",
		},
	})

	if err != nil {
		return nil, err
	}

	var email models.LinkedinEmail
	err = request.JSON(&email)

	user_info.Email = email.Email

	provider.isVerifiedUser = true

	return &user_info, nil
}

/*
	Returns true if the user has a verified email
*/
func (provider *LinkedInOAuthProvider) IsVerifiedUser() bool {
	return provider.isVerifiedUser
}
