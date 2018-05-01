package services

import (
	"net/http"
	"../config"
	"github.com/ASankaran/arbor"
)

const TestURL = config.TestURL

const TestFormat string = "JSON"

var TestRoutes = arbor.RouteCollection {
	arbor.Route {
		"UserAuth",
		"POST",
		"/test/userauth/",
		UserAuth,
	},
	arbor.Route {
		"AdminAuth",
		"POST",
		"/test/adminauth/",
		AdminAuth,
	},
	arbor.Route {
		"NoAuth",
		"POST",
		"/test/noauth/",
		NoAuth,
	},
}

func UserAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r, []string{"User"})
}

func AdminAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r, []string{"Admin"})
}

func NoAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r, nil)
}

