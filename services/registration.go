package services

import (
	"github.com/HackIllinois/api-gateway/config"
	"github.com/HackIllinois/api-gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var RegistrationURL = config.REGISTRATION_SERVICE

const RegistrationFormat string = "JSON"

var RegistrationRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUserRegistration",
		"GET",
		"/registration/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentUserRegistration",
		"POST",
		"/registration/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(CreateRegistration).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentUserRegistration",
		"PUT",
		"/registration/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Applicant"})).ThenFunc(UpdateRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredUserRegistrations",
		"GET",
		"/registration/filter/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetCurrentMentorRegistration",
		"GET",
		"/registration/mentor/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Mentor"})).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentMentorRegistration",
		"POST",
		"/registration/mentor/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Mentor"})).ThenFunc(CreateRegistration).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentMentorRegistration",
		"PUT",
		"/registration/mentor/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Mentor"})).ThenFunc(UpdateRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetUserRegistration",
		"GET",
		"/registration/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetMentorRegistration",
		"GET",
		"/registration/mentor/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetRegistration).ServeHTTP,
	},
}

func GetRegistration(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RegistrationURL+r.URL.String(), RegistrationFormat, "", r)
}

func CreateRegistration(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RegistrationURL+r.URL.String(), RegistrationFormat, "", r)
}

func UpdateRegistration(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, RegistrationURL+r.URL.String(), RegistrationFormat, "", r)
}
