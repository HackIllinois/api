package services

import (
	"github.com/HackIllinois/api-gateway/config"
	"github.com/HackIllinois/api-gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

const DecisionURL = config.DecisionURL

const DecisionFormat string = "JSON"

var DecisionRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentDecision",
		"GET",
		"/decision/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Applicant"})).ThenFunc(GetCurrentDecision).ServeHTTP,
	},
	arbor.Route{
		"GetDecision",
		"GET",
		"/decision/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetDecision).ServeHTTP,
	},
	arbor.Route{
		"UpdateDecision",
		"POST",
		"/decision/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(UpdateDecision).ServeHTTP,
	},
}

func GetCurrentDecision(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, DecisionURL+r.URL.String(), DecisionFormat, "", r)
}

func GetDecision(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, DecisionURL+r.URL.String(), DecisionFormat, "", r)
}

func UpdateDecision(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, DecisionURL+r.URL.String(), DecisionFormat, "", r)
}
