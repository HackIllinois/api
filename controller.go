package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"./models"
	"./errors"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(Authorize)).Methods("GET")
	router.Handle("/code/", alice.New().ThenFunc(Login)).Methods("POST")
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	redirect_url, err := GetAuthorizeRedirect("github")

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	http.Redirect(w, r, redirect_url, 302);
}

func Login(w http.ResponseWriter, r *http.Request) {
	var oauth_code models.OauthCode
	json.NewDecoder(r.Body).Decode(&oauth_code)

	oauth_token, err := GetOauthToken(oauth_code.Code, "github")

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	email, err := GetEmail(oauth_token, "github")

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	// TODO: Get User ID from User Service

	// TODO: Get Roles from DB

	signed_token, err := MakeToken(0, email, []string{"User"})

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	token := models.Token {
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(token)
}
