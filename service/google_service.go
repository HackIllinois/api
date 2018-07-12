package service

import (
	"errors"
	"github.com/HackIllinois/api-auth/config"
	"github.com/HackIllinois/api-auth/models"
	"github.com/levigross/grequests"
)

/*
	Uses a valid oauth token to get the user's primary email
*/
func GetGoogleEmail(oauth_token string) (string, error) {
	request, err := grequests.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json", &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + oauth_token},
	})

	if err != nil {
		return "", err
	}

	var user_info models.GoogleUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.Email == "" {
		return "", errors.New("Invalid oauth token")
	}

	return user_info.Email, nil
}

/*
	Uses a valid oauth code to get a valid oauth token for the user
*/
func GetGoogleOauthToken(code string) (string, error) {
	request, err := grequests.Post("https://www.googleapis.com/oauth2/v4/token", &grequests.RequestOptions{
		Params: map[string]string{
			"client_id":     config.GOOGLE_CLIENT_ID,
			"client_secret": config.GOOGLE_CLIENT_SECRET,
			"code":          code,
			"redirect_uri":  config.AUTH_REDIRECT_URI,
			"grant_type":    "authorization_code",
		},
		Headers: map[string]string{
			"Accept": "application/json",
		},
	})

	if err != nil {
		return "", err
	}

	var oauth_token models.GoogleOauthToken
	err = request.JSON(&oauth_token)

	if err != nil {
		return "", err
	}

	if oauth_token.Token == "" {
		return "", errors.New("Invalid oauth code")
	}

	return oauth_token.Token, nil
}

/*
	Uses a valid oauth token to get the user's unique id
*/
func GetGoogleUniqueId(oauth_token string) (string, error) {
	request, err := grequests.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json", &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + oauth_token},
	})

	if err != nil {
		return "", err
	}

	var user_info models.GoogleUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.ID == "" {
		return "", errors.New("Invalid oauth token")
	}

	return "google" + user_info.ID, nil
}

/*
	Uses a valid oauth token to get the user's username
*/
func GetGoogleUsername(oauth_token string) (string, error) {
	request, err := grequests.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json", &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + oauth_token},
	})

	if err != nil {
		return "", err
	}

	var user_info models.GoogleUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.Name == "" {
		return "", errors.New("Invalid oauth token")
	}

	return user_info.Name, nil
}

/*
	Uses a valid oauth token to get the user's first name
*/
func GetGoogleFirstName(oauth_token string) (string, error) {
	request, err := grequests.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json", &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + oauth_token},
	})

	if err != nil {
		return "", err
	}

	var user_info models.GoogleUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.FirstName == "" {
		return "", errors.New("Invalid oauth token")
	}

	return user_info.FirstName, nil
}

/*
	Uses a valid oauth token to get the user's last name
*/
func GetGoogleLastName(oauth_token string) (string, error) {
	request, err := grequests.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json", &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + oauth_token},
	})

	if err != nil {
		return "", err
	}

	var user_info models.GoogleUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.LastName == "" {
		return "", errors.New("Invalid oauth token")
	}

	return user_info.LastName, nil
}
