package services

import (
	"net/http"

	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const EventFormat string = "JSON"

var EventRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetEventFavorites",
		"GET",
		"/event/favorite/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetEventFavorites).ServeHTTP,
	},
	arbor.Route{
		"AddEventFavorite",
		"POST",
		"/event/favorite/add/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.UserRole}), middleware.IdentificationMiddleware).ThenFunc(AddEventFavorite).ServeHTTP,
	},
	arbor.Route{
		"RemoveEventFavorite",
		"POST",
		"/event/favorite/remove/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.UserRole}), middleware.IdentificationMiddleware).ThenFunc(RemoveEventFavorite).ServeHTTP,
	},
	arbor.Route{
		"MarkUserAsAttendingEvent",
		"POST",
		"/event/track/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(MarkUserAsAttendingEvent).ServeHTTP,
	},
	arbor.Route{
		"GetEventTrackingInfo",
		"GET",
		"/event/track/event/{name}/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetEventTrackingInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUserTrackingInfo",
		"GET",
		"/event/track/user/{id}/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetUserTrackingInfo).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredEvents",
		"GET",
		"/event/filter/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetFilteredEvents).ServeHTTP,
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
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(DeleteEvent).ServeHTTP,
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
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(CreateEvent).ServeHTTP,
	},
	arbor.Route{
		"UpdateEvent",
		"PUT",
		"/event/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateEvent).ServeHTTP,
	},
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func GetFilteredEvents(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func MarkUserAsAttendingEvent(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func GetEventTrackingInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func GetUserTrackingInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func GetEventFavorites(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func AddEventFavorite(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func RemoveEventFavorite(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}
