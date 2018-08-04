package services

import (
	"fmt"
	"github.com/arbor-dev/arbor"
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
	return Routes
}
