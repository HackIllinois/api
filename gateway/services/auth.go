package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const AuthFormat string = "JSON"

var AuthRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUserRoles",
		"GET",
		"/auth/roles/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetUserRoles).ServeHTTP,
	},
	arbor.Route{
		"GetRolesLists",
		"GET",
		"/auth/roles/list/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetRolesLists).ServeHTTP,
	},
	arbor.Route{
		"GetUserListByRole",
		"GET",
		"/auth/roles/list/{role}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetUserListByRole).ServeHTTP,
	},
	arbor.Route{
		"OauthRedirect",
		"GET",
		"/auth/{provider}/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(OauthRedirect).ServeHTTP,
	},
	arbor.Route{
		"OauthCode",
		"POST",
		"/auth/code/{provider}/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(OauthCode).ServeHTTP,
	},
	arbor.Route{
		"GetUserRoles",
		"GET",
		"/auth/roles/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetUserRoles).ServeHTTP,
	},
	arbor.Route{
		"AddUserRole",
		"PUT",
		"/auth/roles/add/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(AddUserRole).ServeHTTP,
	},
	arbor.Route{
		"RemoveUserRole",
		"PUT",
		"/auth/roles/remove/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(RemoveUserRole).ServeHTTP,
	},
	arbor.Route{
		"RefreshToken",
		"GET",
		"/auth/token/refresh/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(RefreshToken).ServeHTTP,
	},
}

func OauthRedirect(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.AUTH_SERVICE+r.URL.String(), AuthFormat, "", r)
}

func OauthCode(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.AUTH_SERVICE+r.URL.String(), AuthFormat, "", r)
}

func GetUserRoles(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.AUTH_SERVICE+r.URL.String(), AuthFormat, "", r)
}

func GetRolesLists(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.AUTH_SERVICE+r.URL.String(), AuthFormat, "", r)
}

func GetUserListByRole(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.AUTH_SERVICE+r.URL.String(), AuthFormat, "", r)
}

func AddUserRole(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.AUTH_SERVICE+r.URL.String(), AuthFormat, "", r)
}

func RemoveUserRole(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.AUTH_SERVICE+r.URL.String(), AuthFormat, "", r)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.AUTH_SERVICE+r.URL.String(), AuthFormat, "", r)
}
