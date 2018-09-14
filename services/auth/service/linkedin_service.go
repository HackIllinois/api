package service

import (
	"errors"
	"fmt"
	"github.com/pattyjogal/api/services/auth/config"
	"github.com/pattyjogal/api/services/auth/models"
	"github.com/levigross/grequests"
)

const LINKEDIN_USER_INFO_URL = "https://api.linkedin.com/v1/people/~:(id,formatted-name,email-address,first-name,last-name)"

/*
	Uses a valid OAuth token to get the user's primary email.
*/
func GetLinkedinEmail(oauth_token string) (string, error) {
	request, err := grequests.Get(LINKEDIN_USER_INFO_URL, &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", oauth_token),
			"Content-Type":  "application/json",
			"x-li-format":   "json"},
	})

	if err != nil {
		return "", err
	}

	var user_info models.LinkedinUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.ID == "" {
		return "", errors.New("Invalid OAuth token.")
	}

	return user_info.Email, nil
}

/*
	Uses a valid OAuth authorization code to get a valid OAuth token for the user.
*/
func GetLinkedinOauthToken(code string, redirect_uri string) (string, error) {
	const LINKEDIN_OAUTH_TOKEN_URL = "https://www.linkedin.com/oauth/v2/accessToken"
	request, err := grequests.Post(LINKEDIN_OAUTH_TOKEN_URL, &grequests.RequestOptions{
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
		return "", err
	}

	var oauth_token models.LinkedinOauthToken
	err = request.JSON(&oauth_token)

	if err != nil {
		return "", err
	}

	if oauth_token.Token == "" {
		return "", errors.New("Invalid OAuth code.")
	}

	return oauth_token.Token, nil
}

/*
	Uses a valid OAuth token to get the user's unique id.
*/
func GetLinkedinUniqueId(oauth_token string) (string, error) {
	request, err := grequests.Get(LINKEDIN_USER_INFO_URL, &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", oauth_token),
			"Content-Type":  "application/json",
			"x-li-format":   "json"},
	})

	if err != nil {
		return "", err
	}

	var user_info models.LinkedinUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.ID == "" {
		return "", errors.New("Invalid OAuth token.")
	}

	return fmt.Sprintf("linkedin%v", user_info.ID), nil
}

/*
	Uses a valid OAuth token to get the user's username.
*/
func GetLinkedinUsername(oauth_token string) (string, error) {
	request, err := grequests.Get(LINKEDIN_USER_INFO_URL, &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", oauth_token),
			"Content-Type":  "application/json",
			"x-li-format":   "json"},
	})

	if err != nil {
		return "", err
	}

	var user_info models.LinkedinUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.ID == "" {
		return "", errors.New("Invalid oauth token")
	}

	return user_info.Username, nil
}

/*
	Uses a valid OAuth token to get the user's name.
*/
func GetLinkedinFirstName(oauth_token string) (string, error) {
	request, err := grequests.Get(LINKEDIN_USER_INFO_URL, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": fmt.Sprintf("Bearer %v", oauth_token),
			"Content-Type": "application/json",
			"x-li-format":  "json"},
	})

	if err != nil {
		return "", err
	}

	var user_info models.LinkedinUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.ID == "" {
		return "", errors.New("Invalid OAuth token.")
	}

	return user_info.FirstName, nil
}

/*
	Uses a valid OAuth token to get the user's last name.
*/
func GetLinkedinLastName(oauth_token string) (string, error) {
	request, err := grequests.Get(LINKEDIN_USER_INFO_URL, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": fmt.Sprintf("Bearer %v", oauth_token),
			"Content-Type": "application/json",
			"x-li-format":  "json"},
	})

	if err != nil {
		return "", err
	}

	var user_info models.LinkedinUserInfo
	err = request.JSON(&user_info)

	if err != nil {
		return "", err
	}

	if user_info.ID == "" {
		return "", errors.New("Invalid OAuth token.")
	}

	return user_info.LastName, nil
}
