package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/HackIllinois/api/services/auth/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{provider}/", alice.New().ThenFunc(Authorize)).Methods("GET")
	router.Handle("/code/{provider}/", alice.New().ThenFunc(Login)).Methods("POST")
	router.Handle("/roles/{id}/", alice.New().ThenFunc(GetRoles)).Methods("GET")
	router.Handle("/roles/add/", alice.New().ThenFunc(AddRole)).Methods("PUT")
	router.Handle("/roles/remove/", alice.New().ThenFunc(RemoveRole)).Methods("PUT")
	router.Handle("/token/refresh/", alice.New().ThenFunc(RefreshToken)).Methods("GET")
}

/*
	Redirects the client to the OAuth authorization url of the specified provider.
*/
func Authorize(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]

	client_application_url := r.URL.Query().Get("redirect_uri")

	if client_application_url == "" {
		client_application_url = config.AUTH_REDIRECT_URI
	}

	oauth_authorization_url, err := service.GetAuthorizeRedirect(provider, client_application_url)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not retrieve OAuth provider authorization code URL."))
	}

	http.Redirect(w, r, oauth_authorization_url, 302)
}

/*
	Converts a valid OAuth authorization code in the request body to an OAuth token.
	Gets basic user information from the OAuth provider and returns a JWT.
*/
func Login(w http.ResponseWriter, r *http.Request) {
	var oauth_code models.OauthCode
	json.NewDecoder(r.Body).Decode(&oauth_code)

	provider := mux.Vars(r)["provider"]

	client_application_url := r.URL.Query().Get("redirect_uri")

	if client_application_url == "" {
		client_application_url = config.AUTH_REDIRECT_URI
	}

	oauth_token, err := service.GetOauthToken(oauth_code.Code, provider, client_application_url)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not get OAuth token.\n" + err.Error()))
	}

	email, err := service.GetEmail(oauth_token, provider)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not fetch user's email from OAuth provider.\n" + err.Error()))
	}

	id, err := service.GetUniqueId(oauth_token, provider)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not fetch user's unique ID from OAuth provider.\n" + err.Error()))
	}

	roles, err := service.GetUserRoles(id, true)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not fetch user's API roles.\n" + err.Error()))
	}

	signed_token, err := service.MakeToken(id, email, roles)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not create HackIllinois API JWT for user."))
	}

	username, err := service.GetUsername(oauth_token, provider)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not fetch user's username from OAuth provider."))
	}

	first_name, err := service.GetFirstName(oauth_token, provider)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not get user's first name from OAuth provider."))
	}

	last_name, err := service.GetLastName(oauth_token, provider)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not get user's last name from OAuth provider."))
	}

	err = service.SendUserInfo(id, username, first_name, last_name, email)

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not send user information to user service."))
	}

	token := models.Token{
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(token)
}

/*
	Gets the roles for the user with the given id.
*/
func GetRoles(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		panic(errors.MALFORMED_REQUEST_ERROR("Must provide ID parameter in request."))
	}

	roles, err := service.GetUserRoles(id, false)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not get user's roles."))
	}

	user_roles := models.UserRoles{
		ID:    id,
		Roles: roles,
	}

	json.NewEncoder(w).Encode(user_roles)
}

/*
	Adds a role to the user with the given id.
*/
func AddRole(w http.ResponseWriter, r *http.Request) {
	var role_modification models.UserRoleModification
	json.NewDecoder(r.Body).Decode(&role_modification)

	if role_modification.ID == "" {
		panic(errors.MALFORMED_REQUEST_ERROR("Must provide ID parameter in request."))
	}

	err := service.AddUserRole(role_modification.ID, role_modification.Role)

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not add user role.\n" + err.Error()))
	}

	roles, err := service.GetUserRoles(role_modification.ID, false)

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not get user's roles.\n" + err.Error()))
	}

	updated_roles := models.UserRoles{
		ID:    role_modification.ID,
		Roles: roles,
	}

	json.NewEncoder(w).Encode(updated_roles)
}

/*
	Removes a role for the user with the given id.
*/
func RemoveRole(w http.ResponseWriter, r *http.Request) {
	var role_modification models.UserRoleModification
	json.NewDecoder(r.Body).Decode(&role_modification)

	if role_modification.ID == "" {
		panic(errors.MALFORMED_REQUEST_ERROR("Must provide ID parameter in request."))
	}

	err := service.RemoveUserRole(role_modification.ID, role_modification.Role)

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not remove user's user role.\n" + err.Error()))
	}

	roles, err := service.GetUserRoles(role_modification.ID, false)

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not fetch user's roles.\n" + err.Error()))
	}

	updated_roles := models.UserRoles{
		ID:    role_modification.ID,
		Roles: roles,
	}

	json.NewEncoder(w).Encode(updated_roles)
}

/*
	Responds with a new JWT token for the user, with updated information.
*/
func RefreshToken(w http.ResponseWriter, r *http.Request) {

	// Fetch user ID from the Identification middleware, and email using the user service

	id := r.Header.Get("HackIllinois-Identity")

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not fetch user info."))
	}

	email := user_info.Email

	// Get the roles from the given user ID

	roles, err := service.GetUserRoles(id, false)

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not fetch user roles.\n" + err.Error()))
	}

	// Create the new token using user ID, email, and (updated) roles.

	signed_token, err := service.MakeToken(id, email, roles)

	if err != nil {
		panic(errors.AUTHORIZATION_ERROR("Could not make a new JWT for the user."))
	}

	new_token := models.Token{
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(new_token)
}
