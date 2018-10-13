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
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(GetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"SetUserInfo",
		"POST",
		"/user/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(SetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"GetCurrentQrCodeInfo",
		"GET",
		"/user/qr/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Attendee"})).ThenFunc(GetCurrentQrCodeInfo).ServeHTTP,
	},
	arbor.Route{
		"GetQrCodeInfo",
		"GET",
		"/user/qr/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetQrCodeInfo).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredUserInfo",
		"GET",
		"/user/filter/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUserInfo",
		"GET",
		"/user/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetUserInfo).ServeHTTP,
	},
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UserURL+r.URL.String(), UserFormat, "", r)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, UserURL+r.URL.String(), UserFormat, "", r)
}

func GetCurrentQrCodeInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}

func GetQrCodeInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}
