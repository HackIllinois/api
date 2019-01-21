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

	oauth_provider, err := service.GetOAuthProvider(provider)

	if err != nil {
		panic(errors.MalformedRequestError(err.Error(), "Invalid OAuth provider."))
	}

	oauth_authorization_url, err := oauth_provider.GetAuthorizationRedirect(client_application_url)

	if err != nil {
		panic(errors.AuthorizationError(err.Error(), "Could not retrieve OAuth provider authorization code URL."))
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

	oauth_provider, err := service.GetOAuthProvider(provider)

	if err != nil {
		panic(errors.MalformedRequestError(err.Error(), "Invalid OAuth provider."))
	}

	err = oauth_provider.Authorize(oauth_code.Code, client_application_url)

	if err != nil {
		panic(errors.AuthorizationError(err.Error(), "Could not get OAuth token."))
	}

	user_info, err := oauth_provider.GetUserInfo()

	if err != nil {
		panic(errors.AuthorizationError(err.Error(), "Could not fetch user's info from OAuth provider."))
	}

	roles, err := service.GetUserRoles(user_info.ID, true)

	if err != nil {
		panic(errors.AuthorizationError(err.Error(), "Could not fetch user's API roles."))
	}

	if oauth_provider.IsVerifiedUser() {
		err = service.AddAutomaticRoleGrants(user_info.ID, user_info.Email)

		if err != nil {
			panic(errors.AuthorizationError(err.Error(), "Could not automatically grant roles to user (based on verified email domain)."))
		}

		roles, err = service.GetUserRoles(user_info.ID, false)

		if err != nil {
			panic(errors.AuthorizationError(err.Error(), "Could not determine user roles, after automatic role grants."))
		}
	}

	signed_token, err := service.MakeToken(user_info.ID, user_info.Email, roles)

	if err != nil {
		panic(errors.AuthorizationError(err.Error(), "Could not create HackIllinois API JWT for user."))
	}

	err = service.SendUserInfo(user_info.ID, user_info.Username, user_info.FirstName, user_info.LastName, user_info.Email)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not send user information to user service."))
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
		panic(errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
	}

	roles, err := service.GetUserRoles(id, false)

	if err != nil {
		panic(errors.AuthorizationError(err.Error(), "Could not get user's roles."))
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
		panic(errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
	}

	err := service.AddUserRole(role_modification.ID, role_modification.Role)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not add user role."))
	}

	roles, err := service.GetUserRoles(role_modification.ID, false)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not get user's roles."))
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
		panic(errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
	}

	err := service.RemoveUserRole(role_modification.ID, role_modification.Role)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not remove user's user role."))
	}

	roles, err := service.GetUserRoles(role_modification.ID, false)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not fetch user's roles."))
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
		panic(errors.DatabaseError(err.Error(), "Could not fetch user info."))
	}

	email := user_info.Email

	// Get the roles from the given user ID

	roles, err := service.GetUserRoles(id, false)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not fetch user roles."))
	}

	// Create the new token using user ID, email, and (updated) roles.

	signed_token, err := service.MakeToken(id, email, roles)

	if err != nil {
		panic(errors.AuthorizationError(err.Error(), "Could not make a new JWT for the user."))
	}

	new_token := models.Token{
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(new_token)
}
