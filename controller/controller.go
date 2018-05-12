package controller

import (
	"github.com/HackIllinois/api-auth/service"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/HackIllinois/api-auth/models"
	"github.com/HackIllinois/api-commons/errors"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{provider}/", alice.New().ThenFunc(Authorize)).Methods("GET")
	router.Handle("/code/{provider}/", alice.New().ThenFunc(Login)).Methods("POST")
	router.Handle("/roles/{id}/", alice.New().ThenFunc(GetRoles)).Methods("GET")
	router.Handle("/roles/", alice.New().ThenFunc(SetRoles)).Methods("PUT")
}

/*
	Redirects the client to the oauth authorization url of the specified provider
*/
func Authorize(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]

	redirect_url, err := service.GetAuthorizeRedirect(provider)

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

	provider := mux.Vars(r)["provider"]

	oauth_token, err := service.GetOauthToken(oauth_code.Code, provider)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	email, err := service.GetEmail(oauth_token, provider)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	id, err := service.GetUniqueId(oauth_token, provider)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	roles, err := service.GetUserRoles(id, true)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	signed_token, err := service.MakeToken(id, email, roles)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	// TODO: Make call to user service to update basic user info

	token := models.Token {
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(token)
}

/*
	Gets the roles for the user with the given id
*/
func GetRoles(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		panic(errors.UnprocessableError("Must provide id parameter"))
	}

	roles, err := service.GetUserRoles(id, false)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	user_roles := models.UserRoles {
		ID: id,
		Roles: roles,
	}

	json.NewEncoder(w).Encode(user_roles)
}

/*
	Updated the roles for the user with the given id
*/
func SetRoles(w http.ResponseWriter, r *http.Request) {
	var user_roles models.UserRoles
	json.NewDecoder(r.Body).Decode(&user_roles)

	if user_roles.ID == "" {
		panic(errors.UnprocessableError("Must provide id parameter"))
	}

	err := service.SetUserRoles(user_roles.ID, user_roles.Roles)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	roles, err := service.GetUserRoles(user_roles.ID, false)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_roles := models.UserRoles {
		ID: user_roles.ID,
		Roles: roles,
	}

	json.NewEncoder(w).Encode(updated_roles)
}
