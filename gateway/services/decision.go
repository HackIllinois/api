package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const DecisionFormat string = "JSON"

var DecisionRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentDecision",
		"GET",
		"/decision/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.ApplicantRole}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentDecision).ServeHTTP,
	},
	arbor.Route{
		"UpdateDecision",
		"POST",
		"/decision/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateDecision).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredDecisions",
		"GET",
		"/decision/filter/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetFilteredDecisions).ServeHTTP,
	},
	arbor.Route{
		"FinalizeDecision",
		"POST",
		"/decision/finalize/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(FinalizeDecision).ServeHTTP,
	},
	arbor.Route{
		"GetDecision",
		"GET",
		"/decision/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetDecision).ServeHTTP,
	},
}

func GetCurrentDecision(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.DECISION_SERVICE+r.URL.String(), DecisionFormat, "", r)
}

func GetDecision(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.DECISION_SERVICE+r.URL.String(), DecisionFormat, "", r)
}

func UpdateDecision(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.DECISION_SERVICE+r.URL.String(), DecisionFormat, "", r)
}

func GetFilteredDecisions(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.DECISION_SERVICE+r.URL.String(), DecisionFormat, "", r)
}

func FinalizeDecision(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.DECISION_SERVICE+r.URL.String(), DecisionFormat, "", r)
}
