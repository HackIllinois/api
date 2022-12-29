package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const CheckinFormat string = "JSON"

var CheckinRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentCheckinInfo",
		"GET",
		"/checkin/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AttendeeRole}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentCheckinInfo).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentCheckinInfo",
		"POST",
		"/checkin/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(CreateCurrentCheckinInfo).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentCheckinInfo",
		"PUT",
		"/checkin/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateCurrentCheckinInfo).ServeHTTP,
	},
	arbor.Route{
		"GetAllCheckedInUsers",
		"GET",
		"/checkin/list/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetAllCheckedInUsers).ServeHTTP,
	},
	arbor.Route{
		"GetCheckinInfo",
		"GET",
		"/checkin/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetCheckinInfo).ServeHTTP,
	},
}

func GetCurrentCheckinInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.CHECKIN_SERVICE+r.URL.String(), CheckinFormat, "", r)
}

func CreateCurrentCheckinInfo(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.CHECKIN_SERVICE+r.URL.String(), CheckinFormat, "", r)
}

func UpdateCurrentCheckinInfo(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.CHECKIN_SERVICE+r.URL.String(), CheckinFormat, "", r)
}

func GetCheckinInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.CHECKIN_SERVICE+r.URL.String(), CheckinFormat, "", r)
}

func GetAllCheckedInUsers(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.CHECKIN_SERVICE+r.URL.String(), CheckinFormat, "", r)
}
