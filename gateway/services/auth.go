package services

import (
	"net/http"

	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

var AuthURL = config.AUTH_SERVICE

const AuthFormat string = "JSON"

var AuthRoutes = arbor.RouteCollection{
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
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetUserRoles).ServeHTTP,
	},
	arbor.Route{
		"AddUserRole",
		"PUT",
		"/auth/roles/remove/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(AddUserRole).ServeHTTP,
	},
	arbor.Route{
		"RemoveUserRole",
		"PUT",
		"/auth/roles/add/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(RemoveUserRole).ServeHTTP,
	},
	arbor.Route{
		"RefreshToken",
		"GET",
		"/auth/token/refresh/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(RefreshToken).ServeHTTP,
	},
}

func OauthRedirect(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

func OauthCode(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

func GetUserRoles(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

func AddUserRole(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

func RemoveUserRole(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}
