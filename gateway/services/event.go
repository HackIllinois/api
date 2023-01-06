package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const EventFormat string = "JSON"

var EventRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetEventFavorites",
		"GET",
		"/event/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetEventFavorites).ServeHTTP,
	},
	arbor.Route{
		"AddEventFavorite",
		"POST",
		"/event/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(AddEventFavorite).ServeHTTP,
	},
	arbor.Route{
		"RemoveEventFavorite",
		"DELETE",
		"/event/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(RemoveEventFavorite).ServeHTTP,
	},
	arbor.Route{
		"MarkUserAsAttendingEvent",
		"POST",
		"/event/track/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(MarkUserAsAttendingEvent).ServeHTTP,
	},
	arbor.Route{
		"GetEventTrackingInfo",
		"GET",
		"/event/track/event/{name}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetEventTrackingInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUserTrackingInfo",
		"GET",
		"/event/track/user/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetUserTrackingInfo).ServeHTTP,
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
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(DeleteEvent).ServeHTTP,
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
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(CreateEvent).ServeHTTP,
	},
	arbor.Route{
		"UpdateEvent",
		"PUT",
		"/event/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateEvent).ServeHTTP,
	},
	arbor.Route{
		"GetEventCode",
		"GET",
		"/event/code/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetEventCode).ServeHTTP,
	},
	arbor.Route{
		"UpdateEventCode",
		"PUT",
		"/event/code/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(PutEventCode).ServeHTTP,
	},
	arbor.Route{
		"Checkin",
		"POST",
		"/event/staff/checkin/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(StaffCheckin).ServeHTTP,
	},
	arbor.Route{
		"Checkin",
		"POST",
		"/event/checkin/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole}), middleware.IdentificationMiddleware).ThenFunc(Checkin).ServeHTTP,
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

func GetEventCode(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func PutEventCode(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func StaffCheckin(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}

func Checkin(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
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
	arbor.DELETE(w, config.EVENT_SERVICE+r.URL.String(), EventFormat, "", r)
}
