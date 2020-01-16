package services

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

const UserFormat string = "JSON"

var UserRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUserInfo",
		"GET",
		"/user/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"SetUserInfo",
		"POST",
		"/user/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(SetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"GetCurrentQrCodeInfo",
		"GET",
		"/user/qr/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentQrCodeInfo).ServeHTTP,
	},
	arbor.Route{
		"GetQrCodeInfo",
		"GET",
		"/user/qr/{id}/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetQrCodeInfo).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredUserInfo",
		"GET",
		"/user/filter/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetUserInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUserInfo",
		"GET",
		"/user/{id}/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetUserInfo).ServeHTTP,
	},
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.USER_SERVICE+r.URL.String(), UserFormat, "", r)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.USER_SERVICE+r.URL.String(), UserFormat, "", r)
}

func GetCurrentQrCodeInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.USER_SERVICE+r.URL.String(), UserFormat, "", r)
}

func GetQrCodeInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.USER_SERVICE+r.URL.String(), UserFormat, "", r)
}
