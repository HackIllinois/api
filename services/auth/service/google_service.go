package service

import (
	"errors"
	"fmt"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/levigross/grequests"
)

type GoogleOAuthProvider struct {
	token          string
	isVerifiedUser bool
}

func NewGoogleOAuth() *GoogleOAuthProvider {
	return &GoogleOAuthProvider{
		token:          "",
		isVerifiedUser: false,
	}
}

/*
	Returns the url to redirects to for OAuth authorization
*/
func (provider *GoogleOAuthProvider) GetAuthorizationRedirect(redirect_uri string) (string, error) {
	return ConstructSafeURL("https", "accounts.google.com", "o/oauth2/v2/auth",
		map[string]string{
			"client_id":     config.GOOGLE_CLIENT_ID,
			"scope":         "profile email",
			"response_type": "code",
			"redirect_uri":  redirect_uri,
		},
	)
}

/*
	Exchanges an OAuth code for an OAuth token
*/
func (provider *GoogleOAuthProvider) Authorize(code string, redirect_uri string) error {
	request, err := grequests.Post("https://www.googleapis.com/oauth2/v4/token", &grequests.RequestOptions{
		Params: map[string]string{
			"client_id":     config.GOOGLE_CLIENT_ID,
			"client_secret": config.GOOGLE_CLIENT_SECRET,
			"code":          code,
			"redirect_uri":  redirect_uri,
			"grant_type":    "authorization_code",
		},
		Headers: map[string]string{
			"Accept": "application/json",
		},
	})

	if err != nil {
		return err
	}

	response_status := fmt.Sprintf("%s", request.String())

	var oauth_token models.GoogleOauthToken
	err = request.JSON(&oauth_token)

	if err != nil {
		return err
	}

	if oauth_token.Token == "" {
		return errors.New("Invalid oauth code. Response: " + response_status)
	}

	provider.token = oauth_token.Token

	return nil
}

/*
	Retrieves user info from the OAuth provider
*/
func (provider *GoogleOAuthProvider) GetUserInfo() (*models.UserInfo, error) {
	var user_info models.UserInfo

	request, err := grequests.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json", &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": "Bearer " + provider.token,
		},
	})

	if err != nil {
		return nil, err
	}

	var google_user_info models.GoogleUserInfo
	err = request.JSON(&google_user_info)

	if err != nil {
		return nil, err
	}

	if google_user_info.ID == "" {
		return nil, errors.New("Invalid oauth token")
	}

	user_info.ID = "google" + google_user_info.ID
	user_info.Email = google_user_info.Email
	user_info.Username = google_user_info.Name
	user_info.FirstName = google_user_info.FirstName
	user_info.LastName = google_user_info.LastName

	provider.isVerifiedUser = google_user_info.IsVerified

	return &user_info, nil
}

/*
	Returns true if the user has a verified email
*/
func (provider *GoogleOAuthProvider) IsVerifiedUser() bool {
	return provider.isVerifiedUser
}
