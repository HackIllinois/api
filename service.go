package main

import (
	"time"
	"errors"
	"./models"
	"./config"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/levigross/grequests"
)

var secret []byte

func init() {
	secret = []byte("secret_string")
}

func MakeToken(id int, email string, roles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"id": id,
		"email": email,
		"roles": roles,
	})

	signed_token, err := token.SignedString(secret)

	return signed_token, err
}

func GetGithubEmail(oauth_token string) (string, error) {
	request, err := grequests.Get("https://api.github.com/user/emails", &grequests.RequestOptions {
		Headers: map[string]string {"Authorization" : "token " + oauth_token},
	})

	if err != nil {
		return "", err
	}

	var emails []models.GithubEmail
	err = request.JSON(&emails)

	if err != nil {
		return "", err
	}

	for _, email := range emails {
		if email.IsPrimary {
			return email.Email, nil
		}
	}

	return "", errors.New("No primary email")
}

func GetGithubOauthToken(code string) (string, error) {
	request, err := grequests.Post("https://github.com/login/oauth/access_token", &grequests.RequestOptions {
		Params: map[string]string {
			"client_id" : config.GITHUB_CLIENT_ID,
			"client_secret" : config.GITHUB_CLIENT_SECRET,
			"code" : code,
		},
		Headers: map[string]string {
			"Accept" : "application/json",
		},
	})

	if err != nil {
		return "", err
	}

	var oauth_token models.GithubOauthToken
	err = request.JSON(&oauth_token)

	if err != nil {
		return "", err
	}

	if oauth_token.Token == "" {
		return "", errors.New("Invalid oauth code")
	}

	return oauth_token.Token, nil
}
