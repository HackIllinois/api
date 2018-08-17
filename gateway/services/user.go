package services

import (
	"github.com/ReflectionsProjections/api/gateway/config"
	"github.com/ReflectionsProjections/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var UserURL = config.USER_SERVICE

const UserFormat string = "JSON"

var UserRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUserInfo",
		"GET",
		"/user/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(GetCurrentUserInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUserInfo",
		"GET",
		"/user/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"SetUserInfo",
		"POST",
		"/user/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(SetUserInfo).ServeHTTP,
	},
}

func GetCurrentUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UserURL+r.URL.String(), UserFormat, "", r)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UserURL+r.URL.String(), UserFormat, "", r)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, UserURL+r.URL.String(), UserFormat, "", r)
}
