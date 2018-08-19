package services

import (
	"github.com/ReflectionsProjections/api/gateway/config"
	"github.com/ReflectionsProjections/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var MailURL = config.MAIL_SERVICE

const MailFormat string = "JSON"

var MailRoutes = arbor.RouteCollection{
	arbor.Route{
		"SendMail",
		"POST",
		"/mail/send/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(SendMail).ServeHTTP,
	},
	arbor.Route{
		"SendMailList",
		"POST",
		"/mail/send/list/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(SendMailList).ServeHTTP,
	},
	arbor.Route{
		"CreateMailList",
		"POST",
		"/mail/list/create/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(CreateMailList).ServeHTTP,
	},
	arbor.Route{
		"AddToMailList",
		"POST",
		"/mail/list/add/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(AddToMailList).ServeHTTP,
	},
	arbor.Route{
		"RemoveFromMailList",
		"POST",
		"/mail/list/remove/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(RemoveFromMailList).ServeHTTP,
	},
	arbor.Route{
		"GetMailList",
		"GET",
		"/mail/list/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetMailList).ServeHTTP,
	},
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MailURL+r.URL.String(), MailFormat, "", r)
}

func SendMailList(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MailURL+r.URL.String(), MailFormat, "", r)
}

func CreateMailList(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MailURL+r.URL.String(), MailFormat, "", r)
}

func AddToMailList(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MailURL+r.URL.String(), MailFormat, "", r)
}

func RemoveFromMailList(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MailURL+r.URL.String(), MailFormat, "", r)
}

func GetMailList(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MailURL+r.URL.String(), MailFormat, "", r)
}
