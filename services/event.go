package services

import (
	"net/http"

	"github.com/HackIllinois/api-gateway/config"
	"github.com/HackIllinois/api-gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

var EventURL = config.EVENT_SERVICE

const EventFormat string = "JSON"

var EventRoutes = arbor.RouteCollection{
	arbor.Route{
		"MarkUserAsAttendingEvent",
		"POST",
		"/event/track/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(MarkUserAsAttendingEvent).ServeHTTP,
	},
	arbor.Route{
		"GetEventTrackingInfo",
		"GET",
		"/event/track/event/{name}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetEventTrackingInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUserTrackingInfo",
		"GET",
		"/event/track/user/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetUserTrackingInfo).ServeHTTP,
	},
	arbor.Route{
		"GetEvent",
		"GET",
		"/event/{name}/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetEvent).ServeHTTP,
	},
	arbor.Route{
		"DeleteEvent",
		"DELETE",
		"/event/{name}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(DeleteEvent).ServeHTTP,
	},
	arbor.Route{
		"GetAllEvents",
		"GET",
		"/event/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetEvent).ServeHTTP,
	},
	arbor.Route{
		"CreateEvent",
		"POST",
		"/event/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(CreateEvent).ServeHTTP,
	},
	arbor.Route{
		"UpdateEvent",
		"PUT",
		"/event/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(UpdateEvent).ServeHTTP,
	},
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, EventURL+r.URL.String(), EventFormat, "", r)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, EventURL+r.URL.String(), EventFormat, "", r)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, EventURL+r.URL.String(), EventFormat, "", r)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, EventURL+r.URL.String(), EventFormat, "", r)
}

func MarkUserAsAttendingEvent(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, EventURL+r.URL.String(), EventFormat, "", r)
}

func GetEventTrackingInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, EventURL+r.URL.String(), EventFormat, "", r)
}

func GetUserTrackingInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, EventURL+r.URL.String(), EventFormat, "", r)
}
