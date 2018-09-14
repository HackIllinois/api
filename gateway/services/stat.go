package services

import (
	"net/http"

	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

var StatURL = config.STAT_SERVICE

const StatFormat string = "JSON"

var StatRoutes = arbor.RouteCollection{
	arbor.Route{
		"RegisterService",
		"POST",
		"/stat/service/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(RegisterService).ServeHTTP,
	},
	arbor.Route{
		"GetService",
		"GET",
		"/stat/service/{name}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetService).ServeHTTP,
	},
	arbor.Route{
		"GetStat",
		"GET",
		"/stat/{name}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetStat).ServeHTTP,
	},
	arbor.Route{
		"GetAllStats",
		"GET",
		"/stat/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetAllStats).ServeHTTP,
	},
}

func RegisterService(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, StatURL+r.URL.String(), StatFormat, "", r)
}

func GetService(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, StatURL+r.URL.String(), StatFormat, "", r)
}

func GetStat(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, StatURL+r.URL.String(), StatFormat, "", r)
}

func GetAllStats(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, StatURL+r.URL.String(), StatFormat, "", r)
}
