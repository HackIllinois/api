package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"./models"
	"./config"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(Authorize)).Methods("GET")
	router.Handle("/code/", alice.New().ThenFunc(Login)).Methods("POST")
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/login/oauth/authorize?client_id=" + config.GITHUB_CLIENT_ID, 302);
}

func Login(w http.ResponseWriter, r *http.Request) {
	var oauth_code models.OauthCode
	json.NewDecoder(r.Body).Decode(&oauth_code)

	oauth_token, err := GetOauthToken(oauth_code.Code)

	if err != nil {
		// TODO: Handle error
	}

	email, err := GetGithubEmail(oauth_token)

	if err != nil {
		// TODO: Handle error
	}

	signed_token, err := MakeToken(0, email, []string{"User"})

	if err != nil {
		// TODO: Handle error
	}

	token := models.Token {
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(token)
}
