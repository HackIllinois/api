package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"./models"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/github/", alice.New().ThenFunc(LoginGithub)).Methods("POST")
}

func LoginGithub(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	json.NewDecoder(r.Body).Decode(&login)

	// TODO: Login github here
	email, _ := GetGithubEmail(login.Oauth)

	signed_token, err := MakeToken(0, email, []string{"User"})

	if err != nil {
		// TODO: Handle error
	}

	token := models.Token {
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(token)
}
