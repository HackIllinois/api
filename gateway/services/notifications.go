package services

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"

	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var NotificationsURL = config.NOTIFICATIONS_SERVICE

const NotificationsFormat string = "JSON"

var NotificationsRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetAllTopics",
		"GET",
		"/notifications/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole})).ThenFunc(GetAllTopics).ServeHTTP,
	},
	arbor.Route{
		"GetAllNotifications",
		"GET",
		"/notifications/all/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.UserRole})).ThenFunc(GetAllNotifications).ServeHTTP,
	},
	arbor.Route{
		"CreateTopic",
		"POST",
		"/notifications/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole})).ThenFunc(CreateTopic).ServeHTTP,
	},
	arbor.Route{
		"RegisterDeviceToUser",
		"POST",
		"/notifications/device/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.UserRole})).ThenFunc(RegisterDeviceToUser).ServeHTTP,
	},
	arbor.Route{
		"GetNotificationsForTopic",
		"GET",
		"/notifications/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.UserRole})).ThenFunc(GetNotificationsForTopic).ServeHTTP,
	},
	arbor.Route{
		"DeleteTopic",
		"DELETE",
		"/notifications/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole})).ThenFunc(DeleteTopic).ServeHTTP,
	},
	arbor.Route{
		"PublishNotification",
		"POST",
		"/notifications/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole})).ThenFunc(PublishNotification).ServeHTTP,
	},
	arbor.Route{
		"AddUsersToTopic",
		"POST",
		"/notifications/{id}/add/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole})).ThenFunc(AddUsersToTopic).ServeHTTP,
	},
	arbor.Route{
		"RemoveUsersFromTopic",
		"POST",
		"/notifications/{id}/remove/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole})).ThenFunc(RemoveUsersFromTopic).ServeHTTP,
	},
	arbor.Route{
		"GetTopicInfo",
		"GET",
		"/notifications/{id}/info/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole})).ThenFunc(GetTopicInfo).ServeHTTP,
	},
}

func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func GetTopicInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func PublishNotification(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func GetNotificationsForTopic(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func AddUsersToTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func RemoveUsersFromTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func RegisterDeviceToUser(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}
