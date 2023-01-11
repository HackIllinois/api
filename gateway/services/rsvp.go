package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const RsvpFormat string = "JSON"

var RsvpRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentRsvpInfo",
		"GET",
		"/rsvp/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.ApplicantRole}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentRsvpInfo).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentRsvpInfo",
		"POST",
		"/rsvp/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.ApplicantRole}), middleware.IdentificationMiddleware).ThenFunc(CreateCurrentRsvpInfo).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentRsvpInfo",
		"PUT",
		"/rsvp/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.ApplicantRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateCurrentRsvpInfo).ServeHTTP,
	},
	arbor.Route{
		"GetRsvpInfo",
		"GET",
		"/rsvp/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetRsvpInfo).ServeHTTP,
	},
}

func GetCurrentRsvpInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.RSVP_SERVICE+r.URL.String(), RsvpFormat, "", r)
}

func CreateCurrentRsvpInfo(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.RSVP_SERVICE+r.URL.String(), RsvpFormat, "", r)
}

func UpdateCurrentRsvpInfo(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.RSVP_SERVICE+r.URL.String(), RsvpFormat, "", r)
}

func GetRsvpInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.RSVP_SERVICE+r.URL.String(), RsvpFormat, "", r)
}
