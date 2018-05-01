package services

import (
	"fmt"
	"net/http"
	"github.com/arbor-dev/arbor"
)

var Routes = arbor.RouteCollection {
	arbor.Route {
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
	Routes = append(Routes, TestRoutes...)
	return Routes
}
