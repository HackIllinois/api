package services

import (
	"github.com/HackIllinois/api-gateway/config"
	"github.com/HackIllinois/api-gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

const AuthURL = config.AuthURL

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
		"SetUserRoles",
		"PUT",
		"/auth/roles/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(SetUserRoles).ServeHTTP,
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

func SetUserRoles(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}
