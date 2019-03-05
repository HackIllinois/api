package services

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"

	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

const NotificationsFormat string = "JSON"

var NotificationsRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetAllTopics",
		"GET",
		"/notifications/topic/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole})).ThenFunc(GetAllTopics).ServeHTTP,
	},
	arbor.Route{
		"CreateTopic",
		"POST",
		"/notifications/topic/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole})).ThenFunc(CreateTopic).ServeHTTP,
	},
	arbor.Route{
		"GetAllNotifications",
		"GET",
		"/notifications/topic/all/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.UserRole})).ThenFunc(GetAllNotifications).ServeHTTP,
	},
	arbor.Route{
		"GetAllPublicNotifications",
		"GET",
		"/notifications/topic/public/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetAllPublicNotifications).ServeHTTP,
	},
	arbor.Route{
		"GetNotificationsForTopic",
		"GET",
		"/notifications/topic/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole})).ThenFunc(GetNotificationsForTopic).ServeHTTP,
	},
	arbor.Route{
		"PublishNotificationToTopic",
		"POST",
		"/notifications/topic/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole})).ThenFunc(PublishNotificationToTopic).ServeHTTP,
	},
	arbor.Route{
		"DeleteTopic",
		"DELETE",
		"/notifications/topic/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.AdminRole})).ThenFunc(DeleteTopic).ServeHTTP,
	},
	arbor.Route{
		"SubscribeToTopic",
		"POST",
		"/notifications/topic/{id}/subscribe/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.UserRole})).ThenFunc(SubscribeToTopic).ServeHTTP,
	},
	arbor.Route{
		"UnsubscribeToTopic",
		"POST",
		"/notifications/topic/{id}/unsubscribe/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.UserRole})).ThenFunc(UnsubscribeToTopic).ServeHTTP,
	},
	arbor.Route{
		"RegisterDeviceToUser",
		"POST",
		"/notifications/device/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]models.Role{models.UserRole})).ThenFunc(RegisterDeviceToUser).ServeHTTP,
	},
}

func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func GetAllPublicNotifications(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func GetNotificationsForTopic(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func PublishNotificationToTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func SubscribeToTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func UnsubscribeToTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}

func RegisterDeviceToUser(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.NOTIFICATIONS_SERVICE+r.URL.String(), NotificationsFormat, "", r)
}
