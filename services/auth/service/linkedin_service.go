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
			"scope":         "r_basicprofile r_emailaddress",
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

	request, err := grequests.Get("https://api.linkedin.com/v1/people/~:(id,formatted-name,email-address,first-name,last-name)", &grequests.RequestOptions{
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
	user_info.Username = linkedin_user_info.Username
	user_info.Email = linkedin_user_info.Email
	user_info.FirstName = linkedin_user_info.FirstName
	user_info.LastName = linkedin_user_info.LastName

	provider.isVerifiedUser = true

	return &user_info, nil
}

/*
	Returns true if the user has a verified email
*/
func (provider *LinkedInOAuthProvider) IsVerifiedUser() bool {
	return provider.isVerifiedUser
}
