package services

import (
	"fmt"
	"net/http"

	"github.com/HackIllinois/api/gateway/config"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

var ServiceLocations map[string]string

func Initialize() error {
	ServiceLocations = map[string]string{
		"auth":          config.AUTH_SERVICE,
		"user":          config.USER_SERVICE,
		"registration":  config.REGISTRATION_SERVICE,
		"decision":      config.DECISION_SERVICE,
		"rsvp":          config.RSVP_SERVICE,
		"checkin":       config.CHECKIN_SERVICE,
		"upload":        config.UPLOAD_SERVICE,
		"mail":          config.MAIL_SERVICE,
		"event":         config.EVENT_SERVICE,
		"stat":          config.STAT_SERVICE,
		"notifications": config.NOTIFICATIONS_SERVICE,
		"project":       config.PROJECT_SERVICE,
		"profile":       config.PROFILE_SERVICE,
	}

	return nil
}

var Routes = arbor.RouteCollection{
	arbor.Route{
		"Gateway",
		"GET",
		"/",
		Gateway,
	},
}

func Gateway(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The API Gateway Lives")
}

func RegisterAPIs() arbor.RouteCollection {

	// arbor does not currently handle preflight requests
	// so for now we handle them here
	Routes = append(Routes, arbor.Route{
		"Preflight",
		"OPTIONS",
		"/{name:.*}",
		alice.New().ThenFunc(AllowCorsPreflight).ServeHTTP,
	})

	Routes = append(Routes, AuthRoutes...)
	Routes = append(Routes, UserRoutes...)
	Routes = append(Routes, RegistrationRoutes...)
	Routes = append(Routes, DecisionRoutes...)
	Routes = append(Routes, RsvpRoutes...)
	Routes = append(Routes, CheckinRoutes...)
	Routes = append(Routes, UploadRoutes...)
	Routes = append(Routes, MailRoutes...)
	Routes = append(Routes, EventRoutes...)
	Routes = append(Routes, StatRoutes...)
	Routes = append(Routes, NotificationsRoutes...)
	Routes = append(Routes, ProjectRoutes...)
	Routes = append(Routes, ProfileRoutes...)
	Routes = append(Routes, HealthRoutes...)
	Routes = append(Routes, ReloadRoutes...)
	return Routes
}

func AllowCorsPreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, CONNECT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, Content-Disposition")
	w.WriteHeader(http.StatusOK)
}
