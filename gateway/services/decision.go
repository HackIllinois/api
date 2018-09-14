package services

import (
	"net/http"

	"github.com/pattyjogal/api/gateway/config"
	"github.com/pattyjogal/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

var DecisionURL = config.DECISION_SERVICE

const DecisionFormat string = "JSON"

var DecisionRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentDecision",
		"GET",
		"/decision/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Applicant"})).ThenFunc(GetCurrentDecision).ServeHTTP,
	},
	arbor.Route{
		"UpdateDecision",
		"POST",
		"/decision/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(UpdateDecision).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredDecisions",
		"GET",
		"/decision/filter/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetFilteredDecisions).ServeHTTP,
	},
	arbor.Route{
		"FinalizeDecision",
		"POST",
		"/decision/finalize/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(FinalizeDecision).ServeHTTP,
	},
	arbor.Route{
		"GetDecision",
		"GET",
		"/decision/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetDecision).ServeHTTP,
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

func GetFilteredDecisions(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, DecisionURL+r.URL.String(), DecisionFormat, "", r)
}

func FinalizeDecision(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, DecisionURL+r.URL.String(), DecisionFormat, "", r)
}
