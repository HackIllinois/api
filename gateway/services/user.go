package services

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
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
		alice.New(middleware.AuthMiddleware([]string{"User"}), middleware.IdentificationMiddleware).ThenFunc(GetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"SetUserInfo",
		"POST",
		"/user/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"}), middleware.IdentificationMiddleware).ThenFunc(SetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredUserInfo",
		"GET",
		"/user/filter/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"}), middleware.IdentificationMiddleware).ThenFunc(GetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUserInfo",
		"GET",
		"/user/{id}/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"}), middleware.IdentificationMiddleware).ThenFunc(GetUserInfo).ServeHTTP,
	},
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UserURL+r.URL.String(), UserFormat, "", r)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, UserURL+r.URL.String(), UserFormat, "", r)
}
