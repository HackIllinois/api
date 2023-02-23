package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const RegistrationFormat string = "JSON"

var RegistrationRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetAllCurrentRegistrations",
		"GET",
		"/registration/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetCurrentUserRegistration",
		"GET",
		"/registration/attendee/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{ // Change back to authtoken.UserRole
		"CreateCurrentUserRegistration",
		"POST",
		"/registration/attendee/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(CreateRegistration).ServeHTTP,
	},
	arbor.Route{ // Change back to authtoken.UserRole
		"UpdateCurrentUserRegistration",
		"PUT",
		"/registration/attendee/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredUserRegistrations",
		"GET",
		"/registration/attendee/list/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetCurrentMentorRegistration",
		"GET",
		"/registration/mentor/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentMentorRegistration",
		"POST",
		"/registration/mentor/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(CreateRegistration).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentMentorRegistration",
		"PUT",
		"/registration/mentor/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.MentorRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredMentorRegistrations",
		"GET",
		"/registration/mentor/list/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetUserRegistration",
		"GET",
		"/registration/attendee/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetMentorRegistration",
		"GET",
		"/registration/mentor/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetRegistration).ServeHTTP,
	},
	arbor.Route{
		"GetAllRegistrations",
		"GET",
		"/registration/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetRegistration).ServeHTTP,
	},
}

func GetRegistration(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.REGISTRATION_SERVICE+r.URL.String(), RegistrationFormat, "", r)
}

func CreateRegistration(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.REGISTRATION_SERVICE+r.URL.String(), RegistrationFormat, "", r)
}

func UpdateRegistration(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.REGISTRATION_SERVICE+r.URL.String(), RegistrationFormat, "", r)
}
