package services

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var PrintURI = config.PRINT_SERVICE
const PrintFormat string = "JSON"

var PrintRoutes = arbor.RouteCollection{
	arbor.Route{
		"CreatePrintJob",
		"POST",
		"/print/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(CreatePrintJob).ServeHTTP, // TODO allow staff and above to print badges
	},
}

func CreatePrintJob(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, PrintURI+r.URL.String(), PrintFormat, "", r)
}
