package services

import (
	"net/http"
	"../config"
	"../middleware"
	"github.com/arbor-dev/arbor"
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
	is_authorized, err := middleware.IsAuthorized(r, []string{"User"})
	if err != nil  || !is_authorized {
		w.WriteHeader(403)
		return
	}
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r)
}

func AdminAuth(w http.ResponseWriter, r *http.Request) {
	is_authorized, err := middleware.IsAuthorized(r, []string{"Admin"})
	if err != nil  || !is_authorized {
		w.WriteHeader(403)
		return
	}
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r)
}

func NoAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, TestURL + r.URL.String(), TestFormat, "", r)
}

