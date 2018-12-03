package services

import (
	"fmt"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

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
	return Routes
}

func AllowCorsPreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, CONNECT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, Content-Disposition")
	w.WriteHeader(http.StatusOK)
}
