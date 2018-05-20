package services

import (
	"github.com/HackIllinois/api-gateway/config"
	"github.com/HackIllinois/api-gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var TestURL = config.TEST_SERVICE

const TestFormat string = "JSON"

var TestRoutes = arbor.RouteCollection{
	arbor.Route{
		"UserAuth",
		"POST",
		"/test/userauth/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(UserAuth).ServeHTTP,
	},
	arbor.Route{
		"AdminAuth",
		"POST",
		"/test/adminauth/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(AdminAuth).ServeHTTP,
	},
	arbor.Route{
		"NoAuth",
		"POST",
		"/test/noauth/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(NoAuth).ServeHTTP,
	},
}

func UserAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL+r.URL.String(), TestFormat, "", r)
}

func AdminAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL+r.URL.String(), TestFormat, "", r)
}

func NoAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL+r.URL.String(), TestFormat, "", r)
}
