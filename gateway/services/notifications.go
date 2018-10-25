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
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(GetAllTopics).ServeHTTP,
	},
}

func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, NotificationsURL+r.URL.String(), NotificationsFormat, "", r)
}
