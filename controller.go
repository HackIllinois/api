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

/*
	Redirects the client to the oauth authorization url of the specified provider
*/
func Authorize(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")

	redirect_url, err := GetAuthorizeRedirect(provider)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	http.Redirect(w, r, redirect_url, 302);
}

/*
	Converts a valid oauth code in the request body to an oauth token
	Gets basic user information from the oauth provider and returns a jwt token
*/
func Login(w http.ResponseWriter, r *http.Request) {
	var oauth_code models.OauthCode
	json.NewDecoder(r.Body).Decode(&oauth_code)

	provider := r.URL.Query().Get("provider")

	oauth_token, err := GetOauthToken(oauth_code.Code, provider)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	email, err := GetEmail(oauth_token, provider)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	id, err := GetUniqueId(oauth_token, provider)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	roles, err := GetUserRoles(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	signed_token, err := MakeToken(id, email, roles)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	token := models.Token {
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(token)
}
