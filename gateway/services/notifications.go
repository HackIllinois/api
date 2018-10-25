package services

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
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
		"/notifications/topics/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetAllTopics).ServeHTTP,
	},
	arbor.Route{
		"CreateTopic",
		"POST",
		"/notifications/topics/create/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(CreateTopic).ServeHTTP,
	},
	arbor.Route{
		"DeleteTopic",
		"POST",
		"/notifications/topics/delete/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(DeleteTopic).ServeHTTP,
	},
	arbor.Route{
		"PublishNotification",
		"POST",
		"/notifications/publish/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(PublishNotification).ServeHTTP,
	},
}

func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}

func PublishNotification(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}
