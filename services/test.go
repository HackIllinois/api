package services

import (
	"net/http"
	"../config"
	"../middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const TestURL = config.TestURL

const TestFormat string = "JSON"

var TestRoutes = arbor.RouteCollection {
	arbor.Route {
		"UserAuth",
		"POST",
		"/test/userauth/",
		alice.New(middleware.AuthMiddleware([]string{"User"})).ThenFunc(UserAuth).ServeHTTP,
	},
	arbor.Route {
		"AdminAuth",
		"POST",
		"/test/adminauth/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(AdminAuth).ServeHTTP,
	},
	arbor.Route {
		"NoAuth",
		"POST",
		"/test/noauth/",
		alice.New().ThenFunc(NoAuth).ServeHTTP,
	},
}

func UserAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r)
}

func AdminAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r)
}

func NoAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r)
}

