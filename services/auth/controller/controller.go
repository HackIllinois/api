package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/HackIllinois/api/services/auth/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/roles/", GetCurrentUserRoles, "GET", router)
	metrics.RegisterHandler("/roles/list/", GetRolesLists, "GET", router)
	metrics.RegisterHandler("/roles/list/{role}/", GetUserListByRole, "GET", router)
	metrics.RegisterHandler("/{provider}/", Authorize, "GET", router)
	metrics.RegisterHandler("/code/{provider}/", Login, "POST", router)
	metrics.RegisterHandler("/roles/{id}/", GetRoles, "GET", router)
	metrics.RegisterHandler("/roles/add/", AddRole, "PUT", router)
	metrics.RegisterHandler("/roles/remove/", RemoveRole, "PUT", router)
	metrics.RegisterHandler("/token/refresh/", RefreshToken, "GET", router)
	metrics.RegisterHandler("/internal/stats/", GetStats, "GET", router)
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
		errors.WriteError(w, r, errors.MalformedRequestError(err.Error(), "Invalid OAuth provider."))
		return
	}

	oauth_authorization_url, err := oauth_provider.GetAuthorizationRedirect(client_application_url)

	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not retrieve OAuth provider authorization code URL."))
		return
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
		errors.WriteError(w, r, errors.MalformedRequestError(err.Error(), "Invalid OAuth provider."))
		return
	}

	err = oauth_provider.Authorize(oauth_code.Code, client_application_url)

	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not get OAuth token."))
		return
	}

	user_info, err := oauth_provider.GetUserInfo()

	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not fetch user's info from OAuth provider."))
		return
	}

	roles, err := service.GetUserRoles(user_info.ID, true)

	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not fetch user's API roles."))
		return
	}

	if oauth_provider.IsVerifiedUser() {
		err = service.AddAutomaticRoleGrants(user_info.ID, user_info.Email)

		if err != nil {
			errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not automatically grant roles to user (based on verified email domain)."))
			return
		}

		roles, err = service.GetUserRoles(user_info.ID, false)

		if err != nil {
			errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not determine user roles, after automatic role grants."))
			return
		}
	}

	signed_token, err := service.MakeToken(user_info, roles)

	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not create HackIllinois API JWT for user."))
		return
	}

	err = service.SendUserInfo(user_info)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not send user information to user service."))
		return
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
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
		return
	}

	roles, err := service.GetUserRoles(id, false)

	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not get user's roles."))
		return
	}

	user_roles := models.UserRoles{
		ID:    id,
		Roles: roles,
	}

	json.NewEncoder(w).Encode(user_roles)
}

/*
	Gets the roles for the current user.
*/
func GetCurrentUserRoles(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
		return
	}

	roles, err := service.GetUserRoles(id, false)

	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not get user's roles."))
		return
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
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
		return
	}

	err := service.AddUserRole(role_modification.ID, role_modification.Role)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not add user role."))
		return
	}

	roles, err := service.GetUserRoles(role_modification.ID, false)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not get user's roles."))
		return
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
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
		return
	}

	err := service.RemoveUserRole(role_modification.ID, role_modification.Role)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not remove user's user role."))
		return
	}

	roles, err := service.GetUserRoles(role_modification.ID, false)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not fetch user's roles."))
		return
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
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch user info."))
		return
	}

	// Get the roles from the given user ID

	roles, err := service.GetUserRoles(id, false)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not fetch user roles."))
		return
	}

	// Create the new token using user ID, email, and (updated) roles.

	signed_token, err := service.MakeToken(user_info, roles)

	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not make a new JWT for the user."))
		return
	}

	new_token := models.Token{
		Token: signed_token,
	}

	json.NewEncoder(w).Encode(new_token)
}

/*
	Responds with a list of valid roles
*/
func GetRolesLists(w http.ResponseWriter, r *http.Request) {
	roles := service.GetValidRoles()

	roles_list := models.UserRoleList{
		Roles: roles,
	}

	json.NewEncoder(w).Encode(roles_list)
}

/*
	Responds with a list of user with the requested role
*/
func GetUserListByRole(w http.ResponseWriter, r *http.Request) {
	role := mux.Vars(r)["role"]

	userids, err := service.GetUsersByRole(role)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve list of users with requested role."))
		return
	}

	user_list := models.UserList{
		UserIDs: userids,
	}

	json.NewEncoder(w).Encode(user_list)
}

/*
	Endpoint to get role stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not fetch registration service statistics."))
		return
	}

	json.NewEncoder(w).Encode(stats)
}
