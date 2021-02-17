package services

import (
	"net/http"

	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const ProfileFormat string = "JSON"

var ProfileRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUserProfile",
		"GET",
		"/profile/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole, models.ApplicantRole})).ThenFunc(GetProfile).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentUserProfile",
		"POST",
		"/profile/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole, models.ApplicantRole})).ThenFunc(CreateProfile).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentUserProfile",
		"PUT",
		"/profile/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.ApplicantRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateProfile).ServeHTTP,
	},
	arbor.Route{
		"DeleteCurrentUserProfile",
		"DELETE",
		"/profile/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(DeleteProfile).ServeHTTP,
	},
	arbor.Route{
		"GetAllProfiles",
		"GET",
		"/profile/list/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetAllProfiles).ServeHTTP,
	},
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}
