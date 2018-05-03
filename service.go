package main

import (
	"time"
	"./models"
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
		return "", nil
	}

	for _, email := range emails {
		if email.IsPrimary {
			return email.Email, nil
		}
	}

	return "", nil
}
