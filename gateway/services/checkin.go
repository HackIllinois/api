package services

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var CheckinURL = config.CHECKIN_SERVICE

const CheckinFormat string = "JSON"

var CheckinRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentCheckinInfo",
		"GET",
		"/checkin/",
		alice.New(middleware.AuthMiddleware([]string{"Attendee"}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentCheckinInfo).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentCheckinInfo",
		"POST",
		"/checkin/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"}), middleware.IdentificationMiddleware).ThenFunc(CreateCurrentCheckinInfo).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentCheckinInfo",
		"PUT",
		"/checkin/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"}), middleware.IdentificationMiddleware).ThenFunc(UpdateCurrentCheckinInfo).ServeHTTP,
	},
	arbor.Route{
		"GetAllCheckedInUsers",
		"GET",
		"/checkin/list/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"}), middleware.IdentificationMiddleware).ThenFunc(GetAllCheckedInUsers).ServeHTTP,
	},
	arbor.Route{
		"GetCurrentQrCodeInfo",
		"GET",
		"/checkin/qr/",
		alice.New(middleware.AuthMiddleware([]string{"Attendee"}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentQrCodeInfo).ServeHTTP,
	},
	arbor.Route{
		"GetCheckinInfo",
		"GET",
		"/checkin/{id}/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"}), middleware.IdentificationMiddleware).ThenFunc(GetCheckinInfo).ServeHTTP,
	},
	arbor.Route{
		"GetQrCodeInfo",
		"GET",
		"/checkin/qr/{id}/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"}), middleware.IdentificationMiddleware).ThenFunc(GetQrCodeInfo).ServeHTTP,
	},
}

func GetCurrentCheckinInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}

func CreateCurrentCheckinInfo(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}

func UpdateCurrentCheckinInfo(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}

func GetCheckinInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}

func GetCurrentQrCodeInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}

func GetQrCodeInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}

func GetAllCheckedInUsers(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CheckinURL+r.URL.String(), CheckinFormat, "", r)
}
